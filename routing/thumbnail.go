package routing

import (
	"net/http"

	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
)

// thumbnail handle GET /mediateq/version/thumbnail/{mediaId}.
// This function generate image file's thumbnail based on the mediaId  and
// the queries parameters: width and height. The resized image is then served back.
func (h handler) thumbnail(ctx *gin.Context) {
	mediaId := ctx.Param("mediaId")

	ctx.JSON(http.StatusOK, jsonutil.Response{"mediaId": mediaId})
}
