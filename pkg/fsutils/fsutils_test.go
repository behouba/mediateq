package fsutils

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {

	// Define test cases
	testCases := []struct {
		name         string
		requestBody  io.Reader
		expectedBody string
		expectedHash string
	}{
		{
			name:         "Test case 1",
			requestBody:  strings.NewReader("Hello, Mediateq!"),
			expectedBody: "Hello, Mediateq!",
			expectedHash: "pWJTWXHB9qtt67AWq1pywKE7V-pQ4kayVObWuub7zC4",
		},
		{
			name:         "Test case 1",
			requestBody:  strings.NewReader("Hello, World!"),
			expectedBody: "Hello, World!",
			expectedHash: "3_1gIbsr1bCvZ2KQgJ7DpTGR3YHH9wpLKGiKNiGCmG8",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the function being tested
			body, hash, err := ReadAndHash(tc.requestBody)

			// Verify the results
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.expectedBody, string(body))
			assert.Equal(t, tc.expectedHash, hash)
		})
	}
}

func TestResizeImage(t *testing.T) {

}
