package local

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/behouba/mediateq/storage"
)

func TestWrite(t *testing.T) {

	cfg := storage.Config{ImagesDir: "static/images", AudiosDir: "static/audio", VideosDir: "static/vidoes"}
	storage, err := NewstorageManager(&cfg)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	filePath, err := storage.Write(ctx, []byte(`Hello world`), "cat.txt")
	if err != nil {
		t.Fatal(err)
	}

	os.RemoveAll("static")

	log.Println(filePath)
}
