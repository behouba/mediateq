package stash

import "context"

// File is a representation of stash file.
type File struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Timestamp int64  `json:"timestamp"`

	// AltSizes only for images
	// Alternatives sizes of the image.
	// Map size and url (Example: "240X400": "hppts://example.com/images/2424/62345234.png")
	AltSizes map[string]string `json:"altSizes,omitempty"`
}

// FileDatabase interface represents the file storage operations.
type FileDatabase interface {
	// Save function save a file to the file storage
	Save(ctx context.Context, f *File) error

	// File function retreive a file object with given id from file storage
	File(ctx context.Context, id string) (File, error)
	// Delete delete a file with the given id from file storage
	Delete(ctx context.Context, id string) error
}

type FileStorage interface {
	Write(ctx context.Context, buff []byte) (url string, err error)
	Remove(ctx context.Context, path string) error
}
