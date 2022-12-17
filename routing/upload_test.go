package routing

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: implement test logic
func TestValidateUploadRequest(t *testing.T) {

}

func TestParseRequestBody(t *testing.T) {

	// Define test cases
	testCases := []struct {
		name         string
		requestBody  io.Reader
		maxFileSize  int64
		expectedBody string
		expectedHash string
	}{
		{
			name:         "Request body too large",
			requestBody:  strings.NewReader("Hello, World!"),
			maxFileSize:  5,
			expectedBody: "Hello",
			expectedHash: "47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU",
		},
		{
			name:         "Request body less than allowed maximum",
			requestBody:  strings.NewReader("Hello, World!"),
			maxFileSize:  100,
			expectedBody: "Hello, World!",
			expectedHash: "47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the function being tested
			body, hash, err := parseRequestBody(tc.requestBody, tc.maxFileSize)

			// Verify the results
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.expectedBody, string(body))
			assert.Equal(t, tc.expectedHash, hash)
		})
	}
}
