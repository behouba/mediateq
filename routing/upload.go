package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/config"
	"github.com/behouba/mediateq/database/schema"
	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

type uploadRequest struct {
	Media *mediateq.Media
}

// Validate check that upload request is valid according the server configuration
func (req uploadRequest) Validate(cfg *config.Mediateq) *jsonutil.Error {
	// Check file size
	if req.Media.Size > cfg.MaxFileSizeBytes {
		return jsonutil.MaxFileSizeExcedeedError(
			fmt.Sprintf("The size of is greather than the maximum allowed upload size: %d bytes", cfg.MaxFileSizeBytes),
		)
	}

	// Check if content type of the file
	if !cfg.IsContentTypeAllowed(req.Media.ContentType) {
		return jsonutil.InvalidContentTypeError(
			fmt.Sprintf("content of type %s is not allowed.", req.Media.ContentType),
		)
	}

	return nil
}

func upload(cfg *config.Mediateq, storage mediateq.Storage, db *schema.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req := uploadRequest{
			Media: &mediateq.Media{
				Origin:      ctx.ClientIP(),
				ContentType: mediateq.ContentType(ctx.ContentType()),
				Size:        ctx.Request.ContentLength,
				UploadName:  ctx.Request.FormValue("filename"),
			},
		}

		if jsonErr := req.Validate(cfg); jsonErr != nil {
			ctx.JSON(http.StatusBadRequest, jsonErr)
			return
		}

		log.Println(req)

	}
}
