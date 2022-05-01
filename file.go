package stash

// File is a representation of stash file.
type File struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Timestamp int64  `json:"timestamp"`
}

// FileStorage interface represents the file storage operations.
type FileStorage interface {
	Save()
	File(id string)
	Delete(id string)
}
