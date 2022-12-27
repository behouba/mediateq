package routing

import (
	"net/http"
	"time"

	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// download handle GET /mediateq/version/download/{mediaId}.
// Serve requested file to the client
func (h handler) download(ctx *gin.Context) {
	mediaID := ctx.Param("mediaId")

	media, err := h.db.MediaTable.SelectByID(ctx, mediaID)
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError("file not found"))
		return
	}

	filePath, err := media.GetFilePath(h.config.Storage.UploadPath)
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.UnknownError("can get file path"))
		return
	}

	// Set the Cache-Control and Expires headers
	ctx.Header("Content-type", string(media.ContentType))
	ctx.Header("Cache-Control", "max-age=86400")
	ctx.Header("Expires", time.Now().Add(time.Hour*24).Format(time.RFC1123))

	ctx.File(filePath)
}

// thumbnail handle GET /mediateq/version/thumbnail/{mediaId}?width={width}&height={height}.
// This function also generate image file's thumbnail based on the mediaId  and
// the queries parameters: width (integer, required) and height (integer, required), and crop (boolean, optional)
func (h handler) thumbnail(ctx *gin.Context) {
	mediaID := ctx.Param("mediaId")

	width, height := queryParamToInt(ctx, "width"), queryParamToInt(ctx, "height")

	crop := ctx.GetBool("crop")

	thumbnail, err := h.db.ThumbnailTable.Select(ctx, mediaID, width, height, crop)
	if err != nil {
		// TODO: Generate thumbnail
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError("file not found"))
		return
	}

	filePath, err := thumbnail.GetFilePath(h.config.Storage.UploadPath)
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.UnknownError("can get file path"))
		return
	}

	// Set the Cache-Control and Expires headers
	ctx.Header("Content-type", string(thumbnail.ContentType))
	ctx.Header("Cache-Control", "max-age=86400")
	ctx.Header("Expires", time.Now().Add(time.Hour*24).Format(time.RFC1123))

	ctx.File(filePath)

	// Generate thumbnail in a go routine

	// // Read media file from storage
	// buffer, err := h.storage.Read(ctx, filePath)
	// if err != nil {
	// 	h.logger.WithField("error", err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, jsonutil.UnknownError("can get read file from storage"))
	// 	return
	// }

	// // Check if the file is an image and if an image resize is needed
	// width := queryParamToInt(ctx, "width")

	// if media.IsImage() && width > 0 {
	// 	buffer, _, err = fsutils.ResizeImage(buffer, width, 0)
	// 	if err != nil {
	// 		h.logger.WithField("error", err.Error())
	// 		ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError(err.Error()))
	// 		return
	// 	}
	// }

	// ctx.Writer.Write(buffer)
}
