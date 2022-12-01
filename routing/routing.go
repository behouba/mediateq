package routing

import (
	"github.com/behouba/mediateq"
	"github.com/gin-gonic/gin"
)

const version = "v0"

func Setup(storage mediateq.Storage, db mediateq.Database) {

	router := gin.Default()

	mediateq := router.Group("/mediateq/" + version)

	{
		mediateq.GET("/info", getServerInfo(db))
		mediateq.POST("/upload", upload(storage, db))
	}

}

func getServerInfo(db mediateq.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func upload(storage mediateq.Storage, db mediateq.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
