package routing

import (
	"log"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/config"
	"github.com/behouba/mediateq/database/schema"
	"github.com/gin-gonic/gin"
)

type uploadRequest struct {
	Media *mediateq.Media
}

// Validate check that upload request is valid according the server configuration
func (req uploadRequest) Validate(cfg *config.Mediateq) error {
	// TODO: implementation
	return nil
}

// parseAndValidateUploadRequest parses file upload request by extracting relevant data then validate
// uploadRequest according to the server configuration
func parseAndValidateUploadRequest(ctx *gin.Context, cfg *config.Mediateq) (*uploadRequest, error) {
	req := uploadRequest{
		Media: &mediateq.Media{
			Origin:      ctx.ClientIP(),
			ContentType: mediateq.ContentType(ctx.ContentType()),
			Size:        ctx.Request.ContentLength,
			UploadName:  ctx.Request.FormValue("filename"),
		},
	}

	if err := req.Validate(cfg); err != nil {
		return nil, err
	}

	return &req, nil
}

func upload(cfg *config.Mediateq, storage mediateq.Storage, db *schema.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req, err := parseAndValidateUploadRequest(ctx, cfg)
		if err != nil {
			// TODO: Define json responses and errors structs (maybe a new package)
			return
		}

		log.Println(req)

	}
}
