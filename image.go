package mediateq

// ImageProcessor interface provides image processing methods
type ImageProcessor interface {
	Resize(buff []byte, width, height int) ([]byte, error)
	Rotage(buff []byte, degree int) ([]byte, error)
	Grayscale(buff []byte) ([]byte, error)
}
