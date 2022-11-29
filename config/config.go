package config

import "github.com/behouba/mediateq/storage"

type Config struct {
	Version string `yaml:"version" json:"version"`
	Host    string `yaml:"host" json:"host"`
	UpTime  int64  `yaml:"up_time" json:"up_time"`
	Storage storage.Config
}
