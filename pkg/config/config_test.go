package config

import (
	"errors"
	"os"
	"testing"

	"github.com/behouba/mediateq"
)

func TestParseAndValidate(t *testing.T) {
	testCases := []struct {
		name     string
		config   Config
		expected error
	}{
		{
			name: "invalid domain",
			config: Config{
				Domain: "http://",
				Database: &Database{
					Host:     "",
					Port:     0,
					DBName:   "",
					Username: "",
				},
				Storage: &Storage{},
			},
			expected: errInvalidDomain,
		},
		{
			name: "missing database config",
			config: Config{
				Domain: "http://www.website.com",
			},
			expected: errMissingDBConfig,
		},
		{
			name: "valid config",
			config: Config{
				Domain: "http://localhost:8080",
				Database: &Database{
					Host:     "127.349.39.20",
					Port:     5000,
					DBName:   "mediateq",
					Username: "user",
				},
				Storage: &Storage{},
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.ParseAndValidate()
			if !errors.Is(err, tc.expected) {
				t.Errorf("expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}

func TestIsContentTypeAllowed(t *testing.T) {

	config := Config{
		AllowedContentTypes: []mediateq.ContentType{
			mediateq.ContentTypeBIMP,
			mediateq.ContentTypeGIF,
			mediateq.ContentTypeJPEG,
			mediateq.ContentTypePNG,
		},
	}

	testCases := []struct {
		name        string
		contentType mediateq.ContentType
		expected    bool
	}{
		{
			name:        "mp4 not allowed",
			contentType: "video/mp4",
			expected:    false,
		},
		{
			name:        " mp3 not allowed",
			contentType: "audio/mp3",
			expected:    false,
		},
		{
			name:        "jpeg allowed",
			contentType: "image/jpeg",
			expected:    true,
		},
		{
			name:        "gif allowed",
			contentType: "image/gif",
			expected:    true,
		},
		{
			name:        "png allowed",
			contentType: "image/png",
			expected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := config.IsContentTypeAllowed(tc.contentType)
			if actual != tc.expected {
				t.Errorf("expected error %v, but got %v", tc.expected, actual)
			}
		})
	}
}

func TestLoad(t *testing.T) {

	// Valid config file test
	t.Run("Valid configuration", func(t *testing.T) {
		// Setup
		filename := "mediateq.yaml"
		data := []byte(`
port: 8080

domain: http://localhost:8080

database:
  host: localhost
  port: 5400
  db_name: mediateq
  username: behouba
  password: password

max_file_size_bytes: 10000000

allowed_content_types:
  - image/jpeg
  - image/png
  - image/gif
  `)

		err := os.WriteFile(filename, data, 0644)
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(filename) // clean up

		// Test
		_, err = Load(filename)
		if err != nil {
			t.Fatal(err)
		}
	})

	// Invalid config file test
	t.Run("Invalid configuration", func(t *testing.T) {
		// Setup
		filename := "mediateq.yaml"
		data := []byte(`
port: 8080


max_file_size_bytes: 10000000

allowed_content_types:
  - image/jpeg
  - image/png
  - image/gif
  `)

		err := os.WriteFile(filename, data, 0644)
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(filename) // clean up

		// Test
		_, err = Load(filename)
		if err == nil {
			t.Fatal("invalid configuration didn't return error")
		}
	})

}
