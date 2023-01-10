package fsutils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/h2non/bimg"
)

// ReadAndHash read a file from an io.Reader and
// create the sha256 hash based on the containt of the request body to be used as filename
func ReadAndHash(r io.Reader) (buf []byte, base64Hash string, err error) {
	hasher := sha256.New()

	teeReader := io.TeeReader(r, hasher)

	buf, err = io.ReadAll(teeReader)
	if err != nil {
		return
	}

	base64Hash = base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))

	return
}

// ResizeImage resize image
// TODO:Comment and write test
func ResizeImage(buf []byte, width, height int, crop bool) (outBuf []byte, base64Hash string, err error) {

	resizedBuf, err := bimg.Resize(
		buf, bimg.Options{Width: width, Height: height, Crop: crop},
	)
	if err != nil {
		return
	}

	return ReadAndHash(bytes.NewBuffer(resizedBuf))
}
