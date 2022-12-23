package fileutil

import (
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/h2non/bimg"
)

// ParseRequestBody read request body and
// create the sha256 hash of the request body to be used as filename
func ParseRequestBody(request io.Reader, maxFileSizeBytes int64) (buffer []byte, hash string, err error) {

	body := io.LimitReader(request, maxFileSizeBytes)

	hasher := sha256.New()

	teeReader := io.TeeReader(body, hasher)

	buffer, err = io.ReadAll(teeReader)
	if err != nil {
		return
	}

	hash = base64.RawURLEncoding.EncodeToString(hasher.Sum(nil)[:])

	return
}

// ResizeImage resize imag
// TODO:Comment and write test
func ResizeImage(buffer []byte, width, height int) (resizedBuffer []byte, hash string, size int64, err error) {

	resizedBuffer, err = bimg.Resize(
		buffer, bimg.Options{Width: width, Height: height},
	)
	if err != nil {
		return
	}

	bufferHash := sha256.Sum256(resizedBuffer)

	hash = base64.RawURLEncoding.EncodeToString(bufferHash[:])

	size = int64(len(resizedBuffer))

	return
}
