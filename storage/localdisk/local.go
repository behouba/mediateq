package localdisk

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"time"

	"github.com/behouba/mediateq/pkg/config"
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

// getSubPath return a formatted representation of the current date
// intended to be used as upload subfolders names in the format {year}/{month}
func getSubPath() string {
	t := time.Now()
	return fmt.Sprintf("%d/%02d", t.Year(), t.Month())
}

func (s storage) Write(ctx context.Context, buff []byte, filename string) (filePath string, err error) {

	subPath := getSubPath()

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

	return filePath, nil
}

func (s storage) Remove(ctx context.Context, path string) error {
	return os.Remove(path)
}
