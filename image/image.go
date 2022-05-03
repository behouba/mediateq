package image

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	buffer, err := bimg.Read("image.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(800, 600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width == 800 && size.Height == 600 {
		fmt.Println("The image size is valid")
	}

	bimg.Write("new.jpg", newImage)
}
