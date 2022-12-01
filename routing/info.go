package routing

// serverInfo provides basic status informations about the server
type serverInfo struct {
	Version          string   `json:"version"`
	Host             string   `json:"host"`
	Uptime           string   `json:"uptime"`
	AllowedFileTypes []string `json:"allowed_file_types"`
	Stats            stats    `json:"stats"`
	AutorizedDomains []string `json:"autorized_domains"`
}

// stats provides basic statistiques about the server
type stats struct {
	Images    int `json:"images"`
	Videos    int `json:"vidoes"`
	Documents int `json:"documents"`
	Audios    int `json:"audios"`
	Errors    int `json:"errors"`
}
