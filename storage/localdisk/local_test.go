package localdisk

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/behouba/mediateq/pkg/config"
)

func TestWrite(t *testing.T) {

	cfg := config.Storage{UploadPath: "/tmp"}
	storage, err := New(&cfg)
	if err != nil {
		t.Fatal(err)
	}

	filename := "hello.txt"

	filePath, err := storage.Write(context.Background(), []byte(`Hello world`), filename)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(filePath)

	log.Println(filePath)
}
