package localdisk

import (
	"context"
	"io/fs"
	"os"
	"path"

	"github.com/behouba/mediateq/pkg/config"
	"github.com/behouba/mediateq/pkg/fileutil"
)

type storage struct {
	cfg *config.Storage
}

func New(cfg *config.Storage) (*storage, error) {

	// Create the file upload directory
	if err := os.MkdirAll(cfg.UploadPath, fs.ModePerm); err != nil {
		return nil, err
	}

	return &storage{cfg}, nil
}

func (s storage) Write(ctx context.Context, buff []byte, filename string) (filePath string, err error) {

	subPath := fileutil.GetSubPath()

	if err = os.MkdirAll(path.Join(s.cfg.UploadPath, subPath), fs.ModePerm); err != nil {
		return
	}

	filePath = path.Join(s.cfg.UploadPath, subPath, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return
	}

	defer file.Close()

	_, err = file.Write(buff)
	if err != nil {
		return
	}

	return path.Join(subPath, filename), nil
}

func (s storage) Remove(ctx context.Context, path string) error {
	return os.Remove(path)
}
