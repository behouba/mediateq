package routing

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/behouba/mediateq/pkg/fileutil"
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
		Timestamp:   time.Now().Unix(),
	}

	// Check file size
	if media.Size > cfg.MaxFileSizeBytes {
		return nil, jsonutil.MaxFileSizeExcedeedError(
			fmt.Sprintf("The size of is greather than the maximum allowed upload size: %d bytes", cfg.MaxFileSizeBytes),
		)
	}

	// Check the content type of the file
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

	media, respErr := extractAndValidateMedia(ctx, h.config)
	if respErr != nil {
		ctx.JSON(http.StatusBadRequest, respErr)
		return
	}

	buffer, hash, err := fileutil.ParseRequestBody(ctx.Request.Body, h.config.MaxFileSizeBytes)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	// // Resize if media is an image and here is a defaut image size width
	if media.IsImage() && (h.config.Storage.DefaultImageSize.Width != 0) {
		buffer, hash, media.Size, err = fileutil.ResizeImage(
			buffer,
			h.config.Storage.DefaultImageSize.Width,
			h.config.Storage.DefaultImageSize.Height,
		)
		if err != nil {
			h.logger.WithField("error", err.Error()).Error()
			ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
			return
		}
	}

	// Set file base64 hash as id
	media.ID = hash

	// Check if we can detect actual content type of the file
	acualContentType := mediateq.ContentType(http.DetectContentType(buffer))
	if acualContentType != media.ContentType && acualContentType != "application/octet-stream" {
		media.ContentType = acualContentType

		if !h.config.IsContentTypeAllowed(media.ContentType) {
			ctx.JSON(http.StatusNotAcceptable, jsonutil.InvalidContentTypeError(
				fmt.Sprintf("content of type %s is not allowed.", media.ContentType),
			),
			)
			return
		}
	}

	path, err := h.storage.Write(ctx, buffer, hash)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.UnknownError("failed to write file to storage"))
		return
	}

	media.URL, err = url.JoinPath(h.config.Domain, downloadPath, path)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	// TODO: handle case of duplicate upload
	media.NID, err = h.db.MediaTable.Insert(ctx, media)
	if err != nil {
		h.logger.WithField("database-error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	ctx.JSON(http.StatusOK, jsonutil.Response{"media": media})

}
