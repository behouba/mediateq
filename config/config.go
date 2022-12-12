package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration object of mediateq server
type Mediateq struct {
	Version  string    `yaml:"version" json:"version"`
	Port     int       `yaml:"port"`
	Domain   string    `yaml:"domain" json:"domain"`
	Database *Database `yaml:"database"`
	Storage  *Storage  `yaml:"storage"`
	Image    *File     `yaml:"image"`
	Audio    *File     `yaml:"audio"`
	Video    *File     `yaml:"video"`
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

	return &cfg, nil
}
