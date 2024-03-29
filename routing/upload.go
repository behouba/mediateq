package routing

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/behouba/mediateq/pkg/fsutils"
	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

// extractAndValidateMedia extract data of the media file to be uploaded
// then check that upload request is valid according the server configuration
func extractAndValidateMedia(ctx *gin.Context, cfg *config.Config) (*mediateq.Media, *jsonutil.Error) {

	// Extract data of the media file to be uploaded
	media := &mediateq.Media{
		ID:          ksuid.New().String(),
		Origin:      ctx.ClientIP(),
		ContentType: mediateq.ContentType(ctx.ContentType()),
		SizeBytes:   ctx.Request.ContentLength,
		Timestamp:   time.Now().Unix(),
	}

	// Check file size
	if media.SizeBytes > cfg.MaxFileSizeBytes {
		return nil, jsonutil.MaxFileSizeExcedeedError(
			fmt.Sprintf("The size of is greather than the maximum allowed upload size: %d bytes", cfg.MaxFileSizeBytes),
		)
	}

	// Check the content type of the file
	if !cfg.IsContentTypeAllowed(media.ContentType) && media.ContentType != mediateq.ContentTypeFormData {
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

	// Try to read the file from the request body
	buffer, base64Hash, err := fsutils.ReadAndHash(ctx.Request.Body)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	// Try to read the file from form when the request body is empty
	if len(buffer) == 0 {

		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			h.logger.WithField("error", err.Error()).Error("failed to get form file")
			ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			h.logger.WithField("error", err.Error()).Error("failed to open form file")
			ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
			return
		}

		defer file.Close()

		buffer, base64Hash, err = fsutils.ReadAndHash(file)
		if err != nil {
			h.logger.WithField("error", err.Error()).Error("failed to read form file")
			ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
			return
		}
	}

	// // Resize image if the file is an image and a defaut image size width is greather than 0
	if media.IsImage() && (h.config.DefaultImageSize.Width > 0) {
		buffer, base64Hash, err = fsutils.ResizeImage(
			buffer,
			h.config.DefaultImageSize.Width,
			h.config.DefaultImageSize.Height,
			false,
		)
		if err != nil {
			h.logger.WithField("error", err.Error()).Error("failed to resize image")
			ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
			return
		}
	}

	// Set file base64 hash
	media.Base64Hash = base64Hash
	// Set the file size
	media.SizeBytes = int64(len(buffer))

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

	filePath, err := media.GetFilePath(h.config.Storage.UploadPath)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.UnknownError("failed to get file path"))
		return
	}

	// Write file to storage
	err = h.storage.Write(ctx, buffer, filePath)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.UnknownError("failed to write file to storage"))
		return
	}

	// Create the file URL
	media.URL, err = url.JoinPath(h.config.Domain, apiBasePath, "download", media.ID)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	// Insert media data to database
	if err := h.db.MediaTable.Insert(ctx, media); err != nil {
		h.logger.WithField("database-error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	// Generate thumbnail in a go routine
	if media.IsImage() {
		go func() {
			for _, size := range h.config.ThumbnailSizes {

				tbuf, base64Hash, err := fsutils.ResizeImage(buffer, size.Width, size.Height, size.Crop)
				if err != nil {
					h.logger.WithField("error", err.Error()).Errorf(
						"failed to generate thumbnail (width: %v, height: %v, crop: %v)", size.Width, size.Height, size.Crop,
					)
					continue
				}

				thumbnail := mediateq.Thumbnail{
					Media: mediateq.Media{
						ID:          media.ID,
						Origin:      media.Origin,
						ContentType: media.ContentType,
						Timestamp:   time.Now().Unix(),
						Base64Hash:  base64Hash,
						SizeBytes:   int64(len(tbuf)),
					},
					ThumbnailSize: size,
				}

				thumbnailFilePath, err := thumbnail.GetFilePath(h.config.Storage.UploadPath)
				if err != nil {
					h.logger.WithField("error", err.Error()).Error("failed to get thumbnail file path")
					continue
				}

				if err := h.storage.Write(ctx, tbuf, thumbnailFilePath); err != nil {
					h.logger.WithField("error", err.Error()).Error("failed to write thumbnail to storage")
					continue
				}

				urlPath, err := url.JoinPath(h.config.Domain, apiBasePath, "thumbnail", media.ID)
				if err != nil {
					h.logger.WithField("error", err.Error()).Error("failed to generate thumbnail URL")
					continue
				}

				thumbnail.URL = fmt.Sprintf(urlPath+"?width=%d&height=%d", thumbnail.Width, thumbnail.Height)

				if thumbnail.Crop {
					thumbnail.URL += "&crop=true"
				}

				if err := h.db.ThumbnailTable.Insert(ctx, &thumbnail); err != nil {
					h.logger.WithField("database-error", err.Error()).Error("failed to save thumbnail to database")
				}

			}

		}()
	}

	ctx.JSON(http.StatusOK, jsonutil.Response{"media": media})

}
