package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

// infoHandler handle request from /stash/info
func (m mux) infoHandler(w http.ResponseWriter, r *http.Request) {

	info := m.getServerInfo()

	bs, err := json.Marshal(info)
	if err != nil {
		w.Write([]byte(`Error` + err.Error()))
		return
	}

	w.Write(bs)
}

func (m mux) getServerInfo() serverInfo {

	return serverInfo{
		Version:          "0.0.1",
		Host:             "localhost",
		Uptime:           fmt.Sprintf("%d second(s)", time.Now().Unix()-m.startTimestamp),
		AllowedFileTypes: []string{"images", "videos", "audios", "documents"},
		Stats: stats{
			Images:    134,
			Videos:    0,
			Documents: 23,
			Audios:    0,
			Errors:    5,
		},
		AutorizedDomains: []string{"localhost"},
	}
}
