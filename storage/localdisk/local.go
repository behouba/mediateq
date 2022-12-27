package localdisk

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/pkg/config"
)

type storage struct {
	cfg *config.Storage
}

func New(cfg *config.Storage) (mediateq.Storage, error) {

	var err error

	cfg.UploadPath, err = filepath.Abs(cfg.UploadPath)
	if err != nil {
		return nil, err
	}

	// Create the file upload directory
	if err := os.MkdirAll(cfg.UploadPath, fs.ModePerm); err != nil {
		return nil, err
	}

	return &storage{cfg}, nil
}

func (s storage) Write(ctx context.Context, buff []byte, filePath string) error {

	// Check if file already exist and return nil
	_, err := os.Stat(filePath)
	if os.IsExist(err) {
		return nil
	}

	// Create subdirectories
	if err := os.MkdirAll(filepath.Dir(filePath), fs.ModePerm); err != nil {
		return err
	}

	// Create a new file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(buff)
	if err != nil {
		return err
	}

	return nil
}

// Read implements mediateq.Storage
func (*storage) Read(ctx context.Context, filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s storage) Remove(ctx context.Context, path string) error {
	return os.Remove(path)
}
