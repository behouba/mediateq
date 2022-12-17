package routing

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/behouba/mediateq"
	"github.com/behouba/mediateq/pkg/config"
	"github.com/digitalcore-ci/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	media, respErr := extractAndValidateMedia(ctx, h.config)
	if respErr != nil {
		ctx.JSON(http.StatusBadRequest, respErr)
		return
	}

	buffer, hash, err := parseRequestBody(ctx.Request.Body, h.config.MaxFileSizeBytes, h.logger)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, jsonutil.InternalServerError())
		return
	}

	h.storage.Write(ctx, buffer, hash)

	// TODO: write to local disk or cloud depending on server configuration

	log.Println(media)

}

// parseRequestBody read request body and
// create the sha256 hash of the request body to be used as filename
func parseRequestBody(
	request io.Reader, maxFileSizeBytes int64, logger *logrus.Logger,
) (buffer []byte, hash string, err error) {

	body := io.LimitReader(request, maxFileSizeBytes)

	buffer, err = ioutil.ReadAll(body)
	if err != nil {
		logger.WithField("error", err.Error()).Error("failed to read upload request body")
		return
	}

	hasher := sha256.New()
	io.TeeReader(body, hasher)

	hash = base64.RawURLEncoding.EncodeToString(hasher.Sum(nil)[:])

	return
}
