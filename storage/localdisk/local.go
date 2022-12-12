package localdisk

import (
	"context"
	"io/fs"
	"os"
	"path"

	"github.com/behouba/mediateq/config"
)

type storage struct {
	cfg *config.Storage
}

func Newstorage(cfg *config.Storage) (*storage, error) {

	if err := os.MkdirAll(cfg.ImagePath, fs.ModePerm); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(cfg.AudioPath, fs.ModePerm); err != nil {
		return nil, err
	}

	return &storage{cfg}, nil
}

func (s storage) Write(ctx context.Context, buff []byte, filename string) (filePath string, err error) {

	filePath = path.Join(s.cfg.ImagePath, filename)

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

func (s storage) Remove(ctx context.Context, path string) error {
	return nil
}
