package localdisk

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/behouba/mediateq/pkg/config"
)

func TestWrite(t *testing.T) {

	cfg := config.Storage{UploadPath: "upload/"}
	storage, err := Newstorage(&cfg)
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
