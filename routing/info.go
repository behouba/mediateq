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
