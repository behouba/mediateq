package mediateq

import "context"

// File is a representation of mediateq file.
type File struct {
	ID        int    `json:"id"`
	Type      string `json:"type"` // image, doc, audio, video
	URL       string `json:"url"`
	Timestamp int64  `json:"timestamp"`

	// AltSizes only for images
	// Alternatives sizes of the image.
	// Map size and url (Example: "240X400": "hppts://example.com/images/2424/62345234.png")
	AltSizes map[string]string `json:"altSizes,omitempty"`
}

// Database interface represents the set of database operations.
type Database interface {
	// Save method save a file to the file storage
	Save(ctx context.Context, f *File) error

	// File method retreive a file object with given id from file storage
	File(ctx context.Context, id string) (File, error)
	// Delete method delete a file with the given id from file storage
	Delete(ctx context.Context, id string) error
}

// FileStorage is an abstration of place where files are stored
type FileStorage interface {
	Write(ctx context.Context, buff []byte, filename string) (url string, err error)
	Remove(ctx context.Context, path string) error
}
