package routing

import (
	"errors"
	"net/http"
	"time"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/database/schema"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// apiVersion is a string representation of the current mediateq server version
const apiVersion = "v0"

const apiBasePath = "/mediateq/" + apiVersion

type handler struct {
	config  *config.Config
	storage mediateq.Storage
	db      *schema.Database
	logger  *logrus.Logger
	info    serverInfo
}

func NewHandler(cfg *config.Config, storage mediateq.Storage, db *schema.Database) (http.Handler, error) {
	router := gin.Default()

	AddRoutes(router, cfg, storage, db)

	return router, nil
}

// serverInfo handle GET /info requests
// It returns information about some configuration data of the server
func (h handler) serverInfo(ctx *gin.Context) {
	h.info.Uptime = time.Now().Unix() - h.info.startTime
	ctx.JSON(http.StatusOK, h.info)
}

func AddRoutes(r http.Handler, cfg *config.Config, storage mediateq.Storage, db *schema.Database) error {

	router, ok := (r).(*gin.Engine)
	if !ok {
		return errors.New("can't convert this handler to gin.Engine")
	}

	mediateq := router.Group(apiBasePath)

	// initialization of logrus for logging
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	h := handler{
		config: cfg, storage: storage,
		db: db, logger: logger,
		info: serverInfo{
			Version:             cfg.Version,
			Domain:              cfg.Domain,
			Port:                cfg.Port,
			startTime:           time.Now().Unix(),
			AllowedContentTypes: cfg.AllowedContentTypes,
			Database:            cfg.Database.Type,
			Storage:             cfg.Storage.Type,
		},
	}

	{
		mediateq.GET("/info", h.serverInfo)
		mediateq.POST("/upload", h.upload)
		mediateq.GET("/download/:base64Hash", h.download)
		mediateq.GET("/thumbnail/:base64Hash", h.thumbnail)

		media := mediateq.Group("/media")
		{
			media.GET("", h.getMediaList)
			media.GET("/:base64Hash", h.getMediaByBase64Hash)
			media.DELETE("/:base64Hash", h.deleteMedia)
		}

	}

	return nil
}
