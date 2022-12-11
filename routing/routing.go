package routing

import (
	"net/http"

	"github.com/behouba/mediateq"
	"github.com/gin-gonic/gin"
)

const version = "v0"

func Setup(cfg *mediateq.Config, storage mediateq.Storage, db mediateq.Database) {

	router := gin.Default()

	mediateq := router.Group("/mediateq/" + version)

	{
		mediateq.GET("/info", getServerInfo(cfg, db))
		mediateq.POST("/upload", uploadFile(storage, db))
	}

}

func getServerInfo(cfg *mediateq.Config, db mediateq.Database) gin.HandlerFunc {
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

func uploadFile(storage mediateq.Storage, db mediateq.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
