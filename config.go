package mediateq

// Configuration object of mediateq server
type Config struct {
	Version          string           `yaml:"version" json:"version"`
	Domain           string           `yaml:"domain" json:"domain"`
	MaxFileSize      int              `yaml:"max_file_size" json:"max_file_size"`
	AllowedFileTypes AllowedFileTypes `yaml:"allowed_file_types"`
	StorageDirs      StorageDirs      `yaml:"storage_dirs"`
}

// AllowedFileTypes specify the types of files allowed to be uploaded to mediateq
// Setting `All` field to `true` mean that every type of file is allowed for upload
type AllowedFileTypes struct {
	Image bool `yaml:"image"`
	Audio bool `yaml:"audio"`
	Video bool `yaml:"video"`
	All   bool `yaml:"all"`
}

// StorageDirs specify the name of directory were each type of files will be stored
type StorageDirs struct {
	Image string `yaml:"image"`
	Audio string `yaml:"audio"`
	Video string `yaml:"video"`
	Other string `yaml:"other"`
}
