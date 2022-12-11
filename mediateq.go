package mediateq

import "context"

type fileType string

const (
	FileTypeImage fileType = "image"
	FileTypeAudio fileType = "audio"
	FileTypeVideo fileType = "video"
)

// ImageProcessor interface provides image processing methods
type ImageProcessor interface {
	Resize(buff []byte, width, height int) ([]byte, error)
	Rotage(buff []byte, degree int) ([]byte, error)
	Grayscale(buff []byte) ([]byte, error)
}

// Media is a representation of mediateq file.
type Media struct {
	NID       int    `json:"nid"`       // Numeric id (db primary key)
	ID        string `json:"id"`        // Unique string identifier of the file
	Type      string `json:"type"`      // image, doc, audio, video
	URL       string `json:"url"`       // url to access the file over internet
	Timestamp int64  `json:"timestamp"` // File creation timestamp
	Size      int64  `json:"size"`      // Size of the file
}

// Database interface represents the set of database operations.
type Database interface {
	// Save method save a file to the file storage
	Save(ctx context.Context, m *Media) error

	// Get method retreive a file object with given id from file storage
	Get(ctx context.Context, id string) (Media, error)
	// Delete method delete a file with the given id from file storage
	Delete(ctx context.Context, id string) error
}

// Storage is an abstration of place where files are stored
type Storage interface {
	Write(ctx context.Context, buff []byte, filename string) (path string, err error)
	Remove(ctx context.Context, path string) error
}
