package localdisk

import (
	"context"
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

	filePath := "/tmp/hello.txt"

	err = storage.Write(context.Background(), []byte(`Hello world`), filePath)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(filePath)

	t.Log(filePath)
}
