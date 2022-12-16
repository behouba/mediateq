package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// extractAndValidateMedia extract data of the media file to be uploaded
// then check that upload request is valid according the server configuration
func extractAndValidateMedia(ctx *gin.Context, cfg *config.Config) (*mediateq.Media, *jsonutil.Error) {

	// Extract data of the media file to be uploaded
	media := &mediateq.Media{
		Origin:      ctx.ClientIP(),
		ContentType: mediateq.ContentType(ctx.ContentType()),
		Size:        ctx.Request.ContentLength,
		UploadName:  ctx.Request.FormValue("filename"),
	}

	// Check file size
	if media.Size > cfg.MaxFileSizeBytes {
		return nil, jsonutil.MaxFileSizeExcedeedError(
			fmt.Sprintf("The size of is greather than the maximum allowed upload size: %d bytes", cfg.MaxFileSizeBytes),
		)
	}

	// Check if content type of the file
	if !cfg.IsContentTypeAllowed(media.ContentType) {
		return nil, jsonutil.InvalidContentTypeError(
			fmt.Sprintf("content of type %s is not allowed.", media.ContentType),
		)
	}

	return media, nil
}

// upload implements POST /upload
// This function handle uploading files to the server
func (h handler) upload(ctx *gin.Context) {

	media, err := extractAndValidateMedia(ctx, h.config)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// reqBody := io.LimitReader(ctx.Request.Body, h.config.MaxFileSizeBytes)

	// TODO: write to local disk or cloud depending on server configuration

	log.Println(media)

}
