package config

import (
	"fmt"
	"io/ioutil"

	"github.com/behouba/mediateq"
	"gopkg.in/yaml.v2"
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
)

// Mediateq is the global configuration object of mediateq server
type Mediateq struct {
	Version             string                 `yaml:"version"`
	Port                int                    `yaml:"port"`
	Domain              string                 `yaml:"domain"`
	Database            *Database              `yaml:"database"`
	Storage             *Storage               `yaml:"storage"`
	AllowedContentTypes []mediateq.ContentType `yaml:"allowed_content_types"`
	MaxFileSizeBytes    int64                  `yaml:"max_file_size_bytes"`
	Image               *File                  `yaml:"image"`
	Audio               *File                  `yaml:"audio"`
	Video               *File                  `yaml:"video"`
}

// IsContentTypeAllowed function searches the
// allowed content Types slice for the given content type and
// returns true if it is found, or false if it is not.
func (m *Mediateq) IsContentTypeAllowed(contentType mediateq.ContentType) bool {
	for _, ct := range m.AllowedContentTypes {
		if ct == contentType {
			return true
		}
	}
	return false
}

// File represent configuration data for a given type of file
type File struct {
	Allowed        bool        `yaml:"allowed"`
	MaxFileSize    int         `yaml:"max_file_size"`
	DefaultSize    ImageSize   `yaml:"image_size"`
	StorageDir     string      `yaml:"storage_dir"`
	ThumbnailSizes []ImageSize `yaml:"thumbnail_sizes,omitempty"`
}

// ImageSize represent a size of an image. To preserve aspect ratio of the original image
// only the  width should set (height should be 0)
type ImageSize struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Storage struct {
	ImagePath string
	AudioPath string
}

// Load loads configuration from yaml file
func Load(filename string) (*Mediateq, error) {

	// Read the YAML file into a byte slice
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML into a Config struct
	var cfg Mediateq
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

	return &cfg, nil
}
