package routing

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// download handle GET /mediateq/version/download/{mediaId}.
// Serve requested file to the client
// This function also generate image file's thumbnail based on the mediaId  and
// the queries parameters: width and height.
func (h handler) download(ctx *gin.Context) {
	// mediaId := ctx.Param("mediaId")

	// _, err := h.db.MediaTable.SelectByHash(ctx, mediaId)
	// if err != nil {
	// 	h.logger.WithField("error", err.Error())
	// 	ctx.JSON(http.StatusInternalServerError, jsonutil.NotFoundError(err.Error()))
	// 	return
	// }

	file, err := os.Open(
		filepath.Join(
			h.config.Storage.UploadPath,
			"2022-12/4siOxL16rCSeWvEeGxBAtqMmF04HffW_qg8zWuOh2MY",
		),
	)
	if err != nil {
		h.logger.WithField("error", err.Error())
		ctx.JSON(http.StatusOK, jsonutil.InternalServerError())
		return
	}

	fileBuffer, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusOK, jsonutil.InternalServerError())
		return
	}

	// Set the Cache-Control and Expires headers
	ctx.Header("Cache-Control", "max-age=86400")
	ctx.Header("Expires", time.Now().Add(time.Hour*24).Format(time.RFC1123))

	ctx.Writer.Write(fileBuffer)

}
