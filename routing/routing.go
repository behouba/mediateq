package routing

import (
	"net/http"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/config"
	"github.com/behouba/mediateq/database/schema"
	"github.com/gin-gonic/gin"
)

const version = "v0"

func NewHandler(cfg *config.Mediateq, storage mediateq.Storage, db *schema.Database) (http.Handler, error) {

	router := gin.Default()

	mediateq := router.Group("/mediateq/" + version)

	{
		mediateq.GET("/info", getServerInfo(cfg, db))
		mediateq.POST("/upload", uploadFile(storage, db))
	}

	return router, nil
}

func getServerInfo(cfg *config.Mediateq, db *schema.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, serverInfo{
			Version:          cfg.Version,
			Host:             cfg.Domain,
			Uptime:           "",
			AllowedFileTypes: []string{},
			Stats:            stats{},
			AutorizedDomains: []string{},
		})
	}
}

func uploadFile(storage mediateq.Storage, db *schema.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
