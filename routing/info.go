package routing

import (
	"github.com/behouba/mediateq"
)

// serverInfo provides basic status informations about the server
type serverInfo struct {
	Version             string                 `json:"version"`
	Domain              string                 `json:"domain"`
	Port                int                    `json:"port"`
	Database            mediateq.DBType        `json:"database"`
	Storage             mediateq.StorageType   `json:"storage"`
	AllowedContentTypes []mediateq.ContentType `json:"allowedContentTypes"`
	Uptime              int64                  `json:"uptime"`
	startTime           int64
}
