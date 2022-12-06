package mediateq

// Configuration object of mediateq server
type Config struct {
	Version     string      `yaml:"version" json:"version"`
	Domain      string      `yaml:"domain" json:"domain"`
	MaxFileSize int         `yaml:"max_file_size" json:"max_file_size"`
	ImageConfig ImageConfig `yaml:"image_config"`
	AudioConfig FileConfig  `yaml:"audio_config"`
	VideoConfig FileConfig  `yaml:"video_config"`
}

// FileConfig represent configuration data for a given type of file
type FileConfig struct {
	Allowed     bool      `yaml:"allowed"`
	MaxFileSize int       `yaml:"max_file_size"`
	DefaultSize ImageSize `yaml:"image_size"`
	StorageDir  string    `yaml:"storage_dir"`
}

// ImageConfig represent configuration data specific to image files
type ImageConfig struct {
	FileConfig
	ThumbnailSizes []ImageSize `yaml:"thumbnail_sizes"`
}

// ImageSize represent a size of an image. To preserve aspect ratio of the original image
// only the  width should set (height should be 0)
type ImageSize struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}
