package routing

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// queryParamToInt extract and convert query parameter to integer
func queryParamToInt(ctx *gin.Context, query string) int {
	qInt, _ := strconv.Atoi(ctx.Query(query))
	return qInt
}

// getMediaList handle GET /media requests
// It return a JSON array of media based on offset and limit query parameters
func (h handler) getMediaList(ctx *gin.Context) {
	offset, limit := queryParamToInt(ctx, "offset"), queryParamToInt(ctx, "limit")

	mediaList, err := h.db.MediaTable.SelectList(ctx, offset, limit)
	if err != nil {
		h.logger.WithField("error", err.Error()).Error("failed to select media list")
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	ctx.JSON(http.StatusOK, jsonutil.Response{"mediaList": mediaList})
}

// getMediaByID handle GET /media/{mediaId}
// It retrieve media data from database and send back JSON response
func (h handler) getMediaByID(ctx *gin.Context) {
	mediaID := ctx.Param("mediaId")

	media, err := h.db.MediaTable.SelectByHash(ctx, mediaID)
	if err != nil {
		// Check if the error is just about non existing media id
		if err == sql.ErrNoRows {
			ctx.JSON(
				http.StatusNotFound, jsonutil.NotFoundError(fmt.Sprintf("media with id %s is not found", mediaID)),
			)
			return
		}

		h.logger.WithField("database-error", err.Error()).Error()
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	ctx.JSON(http.StatusOK, jsonutil.Response{"media": media})
}

// deleteMedia handle DELETE /media/{mediaId}
// This handler will delete media from storage and database by media id if allowed by the server
func (h handler) deleteMedia(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsonutil.Response{})

}
