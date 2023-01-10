package mediateq

import (
	"testing"
)

func TestMediaIsImage(t *testing.T) {
	testData := []struct {
		name     string
		media    Media
		expected bool
	}{
		{
			name:     "valid image",
			media:    Media{ContentType: ContentTypeJPEG},
			expected: true,
		},
		{
			name:     "not an image",
			media:    Media{ContentType: "vidoe/mp4"},
			expected: false,
		},
		{
			name:     "not an image",
			media:    Media{ContentType: "audio/mp3"},
			expected: false,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			if tt.media.IsImage() != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, tt.media.IsImage())
			}
		})
	}
}

func TestGetFilePath(t *testing.T) {

	testData := []struct {
		name             string
		media            Media
		uploadPath       string
		expectedFilePath string
		hasError         bool
	}{
		{
			name:             "test case 1",
			media:            Media{Base64Hash: "mediateq", ContentType: "image/png"},
			uploadPath:       "/tmp",
			expectedFilePath: "/tmp/m/e/diateq",
		},
		{
			name:             "test case 2",
			media:            Media{Base64Hash: "qwerty", ContentType: "image/jpeg"},
			uploadPath:       "/tmp",
			expectedFilePath: "/tmp/q/w/erty",
		},
		{
			name:       "test case 3",
			media:      Media{Base64Hash: "", ContentType: "image/png"},
			uploadPath: "/tmp",
			hasError:   true,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {

			actualFilePath, err := tt.media.GetFilePath(tt.uploadPath)
			if err != nil && !tt.hasError {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.expectedFilePath != actualFilePath {
				t.Fatalf("expected %v, got %v", tt.expectedFilePath, actualFilePath)
			}
		})
	}

}
