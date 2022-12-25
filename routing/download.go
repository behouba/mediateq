package routing

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/behouba/mediateq/pkg/fileutil"
	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// download handle GET /mediateq/version/download/{mediaId}.
// Serve requested file to the client
// This function also generate image file's thumbnail based on the mediaId  and
// the queries parameters: width and height.
func (h handler) download(ctx *gin.Context) {
	base64Hash := ctx.Param("base64Hash")

	media, err := h.db.MediaTable.SelectByHash(ctx, base64Hash)
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError(err.Error()))
		return
	}

	file, err := os.Open(filepath.Join(media.FilePath))
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError(err.Error()))
		return
	}

	fileBuffer, err := io.ReadAll(file)
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError(err.Error()))
		return
	}

	// Check if the file is an image and if an image resize is needed
	width := queryParamToInt(ctx, "width")

	if media.IsImage() && width > 0 {
		fileBuffer, _, _, err = fileutil.ResizeImage(fileBuffer, width, 0)
		if err != nil {
			h.logger.WithField("error", err.Error())
			ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError(err.Error()))
			return
		}
	}

	// Set the Cache-Control and Expires headers
	ctx.Header("Content-type", string(media.ContentType))
	ctx.Header("Cache-Control", "max-age=86400")
	ctx.Header("Expires", time.Now().Add(time.Hour*24).Format(time.RFC1123))

	ctx.Writer.Write(fileBuffer)

}
