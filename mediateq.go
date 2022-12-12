package mediateq

import "context"

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeAudio MediaType = "audio"
	MediaTypeVideo MediaType = "video"
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

// Storage is an abstration of place where files are stored
type Storage interface {
	Write(ctx context.Context, buff []byte, filename string) (path string, err error)
	Remove(ctx context.Context, path string) error
}
