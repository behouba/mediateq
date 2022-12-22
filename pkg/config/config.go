package config

import (
	"fmt"
	"os"
	"regexp"

	"github.com/behouba/mediateq"
	"gopkg.in/yaml.v3"
)

const (
	// Default on port number for the server
	defaultServerPort = 8080

	// Default folder where files should be uploaded
	defaultUploadPath = "upload/"

	// Default max file size in bytes is set to 10MB
	maxFileSizeBytes = 10000000
)

var (
	// Lookup map to quickly check content types supported by the current mediateq server
	supportedContentTypes = map[mediateq.ContentType]bool{
		mediateq.ContentTypeJPEG: true,
		mediateq.ContentTypePNG:  true,
		mediateq.ContentTypeGIF:  true,
		mediateq.ContentTypeBIMP: true,
		mediateq.ContentTypeWEBP: true,
	}

	// Regular expression to validate domain urls
	domainRegex = regexp.MustCompile(`^(https?:\/\/)?(www\.)?([a-zA-Z0-9]+\.)*[a-zA-Z0-9]+(:[0-9]+)?$`)
)

// Config is the global configuration object of mediateq server
type Config struct {
	Version             string                 `yaml:"version"`
	Port                int                    `yaml:"port"`   // Port on which the server should run
	Domain              string                 `yaml:"domain"` // URL of the server domain (example: https://example.com or localhost:8080)
	Database            *Database              `yaml:"database"`
	Storage             *Storage               `yaml:"storage"`
	AllowedContentTypes []mediateq.ContentType `yaml:"allowed_content_types"`
	MaxFileSizeBytes    int64                  `yaml:"max_file_size_bytes"`
}

func (c *Config) ParseAndValidate() error {

	if c.Port == 0 {
		c.Port = defaultServerPort
	}

	//  Validate domain url
	if !domainRegex.MatchString(c.Domain) {
		return errInvalidDomain
	}

	// Check if database config is missing an set a empty default value
	if c.Database == nil {
		return errMissingDBConfig
	}

	// Check if Storage configuration is missing and set a default value
	if c.Storage == nil {
		c.Storage = &Storage{UploadPath: defaultUploadPath}
	}

	// Set default value for upload folder if empty
	if c.Storage.UploadPath == "" {
		c.Storage.UploadPath = defaultUploadPath
	}

	// Set max file size to default value if 0
	if c.MaxFileSizeBytes == 0 {
		c.MaxFileSizeBytes = maxFileSizeBytes
	}

	// When allowed content type array is empty set it to all supported type by default
	if len(c.AllowedContentTypes) == 0 {
		ct := []mediateq.ContentType{}
		for k := range supportedContentTypes {
			ct = append(ct, k)
		}
		c.AllowedContentTypes = ct
	}

	return nil
}

// IsContentTypeAllowed function searches the
// allowed content Types slice for the given content type and
// returns true if it is found, or false if it is not.
func (c *Config) IsContentTypeAllowed(contentType mediateq.ContentType) bool {
	for _, ct := range c.AllowedContentTypes {
		if ct == contentType {
			return true
		}
	}
	return false
}

// ImageSize represent a size of an image. To preserve aspect ratio of the original image
// only the  width should set (height should be 0)
type ImageSize struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

type Database struct {
	Type     mediateq.DBType `yaml:"type"`
	Host     string          `yaml:"host"`
	Port     int             `yaml:"port"`
	DBName   string          `yaml:"db_name"`
	Username string          `yaml:"username"`
	Password string          `yaml:"password"`
}

type Storage struct {
	Type             mediateq.StorageType `yaml:"type"`        // The type of storage used by the to read and write files
	UploadPath       string               `yaml:"upload_path"` // Path to the folder were files should be uploaded
	DeleteAllowed    bool                 `yaml:"delete_allowed"`
	DefaultImageSize ImageSize            `yaml:"default_image_size"`
}

// Load loads mediateq configuration from yaml file
func Load(filename string) (*Config, error) {

	// Read the YAML file into a byte slice
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML into a Config struct
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// Check that allowed content types are all supported
	for _, act := range cfg.AllowedContentTypes {
		if supported := supportedContentTypes[act]; !supported {
			return nil, fmt.Errorf("content type %s is not supported by this version of mediateq", act)
		}
	}

	if err := cfg.ParseAndValidate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
