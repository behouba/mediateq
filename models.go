package mediateq

import (
	"context"
)

// ContentType represents HTTP MIME types sent in Content-type header
type ContentType string

// StorageType represents type of storage where media are stored
type StorageType string

// DB type represents the type of database used by the server
type DBType string

const (
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
	NID               int         `json:"nid"`         // Numeric id (db primary key)
	ID                string      `json:"id"`          // Unique string identifier of the file
	URL               string      `json:"url"`         // url to access the file over internet
	Origin            string      `json:"origin"`      // Origin domain of the file
	ContentType       ContentType `json:"contentType"` //
	Size              int64       `json:"size"`        // Size of the file in bytes
	CreationTimestamp uint64      `json:"creationTimestamp"`
	UploadName        string      `json:"uploadName"` // Media file upload name
	Base64Hash        string      `json:"base64Hash"`
}

// Storage is an abstration of place where files are stored
type Storage interface {
	Write(ctx context.Context, buff []byte, filename string) (path string, err error)
	Remove(ctx context.Context, path string) error
}