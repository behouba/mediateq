package mediateq

import (
	"context"
	"strings"
)

// ContentType represents HTTP MIME types sent in Content-type header
type ContentType string

// StorageType represents type of storage where media are stored
type StorageType string

// DB type represents the type of database used by the server
type DBType string

const (
	ContentTypeFormData ContentType = "multipart/form-data"

	// Supported content types for images
	ContentTypeJPEG ContentType = "image/jpeg"
	ContentTypePNG  ContentType = "image/png"
	ContentTypeGIF  ContentType = "image/gif"
	ContentTypeBIMP ContentType = "image/bimg"
	ContentTypeWEBP ContentType = "image/webp"

	// Storage type options
	StorageTypeLocalDisk StorageType = "localdisk"
	StorageTypeS3        StorageType = "s3"

	// Database options
	DBTypePostgres DBType = "postgres"
	DBTypeSQLite   DBType = "sqlite"
)

// ImageProcessor interface provides image processing methods
type ImageProcessor interface {
	Resize(buff []byte, width, height int) ([]byte, error)
	Rotage(buff []byte, degree int) ([]byte, error)
	Grayscale(buff []byte) ([]byte, error)
}

// Media is a representation of mediateq file.
type Media struct {
	ID          int64       `json:"id"`                   // Numeric id (db primary key)
	Base64Hash  string      `json:"base64Hash"`           // Base64 hash of the file used as a unique string identifier
	URL         string      `json:"url"`                  // url to access the file over internet
	FilePath    string      `json:"filePath"`             // relative path to the file from main upload path
	Origin      string      `json:"origin"`               // Origin domain of the file
	ContentType ContentType `json:"contentType"`          //
	Size        int64       `json:"size"`                 // Size of the file in bytes
	Timestamp   int64       `json:"tmestamp"`             // Media creation timestamp
	UploadName  string      `json:"uploadName,omitempty"` // Media file upload name
}

func (m Media) IsImage() bool {
	cts := strings.Split(string(m.ContentType), "/")
	if len(cts) == 0 {
		return false
	}
	return cts[0] == "image"
}

// Storage is an abstration of place where files are stored
type Storage interface {
	// Write method write file buffer to storage and
	// return the relative path and an error if the write operating fail
	Write(ctx context.Context, buff []byte, filename string) (path string, err error)
	// Remove method should remove a file from
	// the storage given the path to the target file
	Remove(ctx context.Context, path string) error
}
