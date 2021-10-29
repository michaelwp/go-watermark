package go_watermark

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
)

// Resize uses go standard image.Image, unsigned int for size width and height
// will return image object and error status
func Resize(imageObject image.Image, width uint, height uint) (
	img image.Image, err error) {

	// validating image object
	if imageObject == nil {
		return nil, errors.New("image object required")
	}
	return resize.Resize(width, height, imageObject, resize.Lanczos3), nil
}
