package routing

import (
	"net/http"
	"time"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/schema"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/gin-gonic/gin"
)

// apiVersion is a string representation of the current mediateq server version
const apiVersion = "v0"

func NewHandler(cfg *config.Config, storage mediateq.Storage, db *schema.Database) (http.Handler, error) {

	router := gin.Default()

	srvInfo := serverInfo{
		Version:   cfg.Version,
		Domain:    cfg.Domain,
		Port:      cfg.Port,
		StartTime: &time.Time{},
	}

	mediateq := router.Group("/mediateq/" + apiVersion)

	{
		mediateq.GET("/info", getServerInfo(&srvInfo, db))
		mediateq.POST("/upload", upload(cfg, storage, db))
	}

	return router, nil
}

func getServerInfo(srvInfo *serverInfo, db *schema.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, srvInfo)
	}
}
