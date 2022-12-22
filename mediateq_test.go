package mediateq

import "testing"

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
