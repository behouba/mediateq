package mediateq

import (
	"context"
	"fmt"
	"path/filepath"
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
	ID          string      `json:"id"`                   // Numeric id (db primary key)
	Base64Hash  string      `json:"base64Hash"`           // Base64 hash of the file used as a unique string identifier
	URL         string      `json:"url"`                  // url to access the file over internet
	Origin      string      `json:"origin"`               // Origin domain of the file
	ContentType ContentType `json:"contentType"`          //
	SizeBytes   int64       `json:"sizeBytes"`            // Size of the file in bytes
	Timestamp   int64       `json:"tmestamp"`             // Media creation timestamp
	UploadName  string      `json:"uploadName,omitempty"` // Media file upload name
}

// IsImage check if a media file is an image base on it content type
func (m Media) IsImage() bool {
	cts := strings.Split(string(m.ContentType), "/")
	if len(cts) == 0 {
		return false
	}
	return cts[0] == "image"
}

// GetFilePath return the path to a media file
// 2 subdirectories are created for more manageable browsing and use the remainder as the file name.
// For example, if Base64Hash is 'qwerty' and content type is 'image/png' the path will be 'q/w/erty.png'.
func (m Media) GetFilePath(uploadPath string) (string, error) {

	if len(m.Base64Hash) < 3 {
		return "", fmt.Errorf("invalid filePath (Base64Hash too short - min 3 characters): %q", m.Base64Hash)
	}
	if len(m.Base64Hash) > 255 {
		return "", fmt.Errorf("invalid filePath (Base64Hash too long - max 255 characters): %q", m.Base64Hash)
	}

	// ext, err := mime.ExtensionsByType(string(m.ContentType))
	// if err != nil || len(ext) == 0 {
	// 	return "", fmt.Errorf("unable to get media extention for content type %v", m.ContentType)
	// }

	filePath, err := filepath.Abs(filepath.Join(
		uploadPath,
		m.Base64Hash[0:1],
		m.Base64Hash[1:2],
		m.Base64Hash[2:],
	))

	// // Append file extension to filePath
	// filePath = filePath + ext[0]

	if err != nil {
		return "", fmt.Errorf("unable to construct filePath: %w", err)
	}

	// check if the base path is a prefix of the full absolute filePath
	// if so, no directory escape has occurred and the filePath is valid
	if !strings.HasPrefix(filePath, uploadPath) {
		return "", fmt.Errorf("invalid filePath (not within uploadPath %v): %v", uploadPath, filePath)
	}

	return filePath, nil

}

// Storage is an abstration of place where files are stored
type Storage interface {
	// Write method write file buffer to storage and
	// return the relative path and an error if the write operating fail
	Write(ctx context.Context, buff []byte, filePath string) error
	// Read method read a file from storage and return
	// the relative path and an error if the read operation fail
	Read(ctx context.Context, filePath string) ([]byte, error)
	// Remove method should remove a file from
	// the storage given the path to the target file
	Remove(ctx context.Context, path string) error
}
