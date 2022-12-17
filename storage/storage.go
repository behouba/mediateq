package storage

import (
	"fmt"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/behouba/mediateq/storage/localdisk"
)

func New(cfg *config.Storage) (mediateq.Storage, error) {

	if cfg.Type == mediateq.StorageTypeLocalDisk {
		return localdisk.New(cfg)
	}

	return nil, fmt.Errorf("storage of type %s is not supported", cfg.Type)
}
