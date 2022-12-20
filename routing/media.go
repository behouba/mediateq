package routing

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// getMediaByID handle GET /media/{mediaId}
// It retrieve media data from database and send back JSON response
func (h handler) getMediaByID(ctx *gin.Context) {
	mediaID := ctx.Param("mediaId")

	media, err := h.db.MediaTable.SelectByID(ctx, mediaID)
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
