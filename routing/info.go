package routing

import (
	"time"

	"github.com/behouba/mediateq"
)

// serverInfo provides basic status informations about the server
type serverInfo struct {
	Version             string                 `json:"version"`
	Domain              string                 `json:"domain"`
	Port                int                    `json:"port"`
	StartTime           time.Time              `json:"startTime"`
	AllowedContentTypes []mediateq.ContentType `json:"allowedContentTypes"`
}

// stats provides basic statistiques about the server
type stats struct {
	Images    int `json:"images"`
	Videos    int `json:"vidoes"`
	Audios    int `json:"audios"`
	Documents int `json:"documents"`
	Errors    int `json:"errors"`
}
