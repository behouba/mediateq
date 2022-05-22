package local

import (
	"context"
	"io/fs"
	"os"
	"path"

	"github.com/behouba/stash/storage"
)

type storageManager struct {
	cfg *storage.Config
}

func NewstorageManager(cfg *storage.Config) (*storageManager, error) {

	if err := os.MkdirAll(cfg.ImagesDir, fs.ModePerm); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(cfg.AudiosDir, fs.ModePerm); err != nil {
		return nil, err
	}

	return &storageManager{cfg}, nil
}

func (s storageManager) Write(ctx context.Context, buff []byte, filename string) (filePath string, err error) {

	filePath = path.Join(s.cfg.ImagesDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return
	}

	defer file.Close()

	_, err = file.Write(buff)
	if err != nil {
		return
	}

	return filePath, nil
}

func (s storageManager) Remove(ctx context.Context, path string) error {
	return nil
}
