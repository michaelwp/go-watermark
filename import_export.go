package go_watermark

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Import get image object from file
// will return image object & error status
func importImg(imageFile string) (
	imageObject image.Image, imageType string, err error) {

	// open file
	f, err := os.Open(imageFile)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	// decode file to image object
	imageObject, imageType, err = image.Decode(f)
	return imageObject, imageType, err
}

// Export image object into file with specific type and file name.
// default filename & image type is `output.jpeg`.
// will return error status.
func exportImg(imageObject image.Image, imageType string, fileName string) (
	err error) {

	var out *os.File

	// validating image object arguments
	if imageObject == nil {
		return errors.New("image object required")
	}

	// transform image type to lower case
	imageType = strings.ToLower(imageType)

	// validating image type
	if imageType == "" || imageType == "jpg" {
		// set to default image type
		imageType = "jpeg"
	}

	// validating file name
	if fileName == "" {
		fileName = "output"
	}

	// create output file
	out, err = os.Create(fmt.Sprintf("%s.%s", fileName, imageType))
	if err != nil {
		return err
	}

	// export image
	switch imageType {
	case "jpeg":
		err = exportJPEG(imageObject, out)
		if err != nil {
			return err
		}
	case "PNG":
		err = exportPNG(imageObject, out)
		if err != nil {
			return err
		}
	case "GIF":
		err = exportGIF(imageObject, out)
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown data type")
	}

	return nil
}

// exportJPEG exporting image object to jpeg/jpg file
// will return error status
func exportJPEG(imageObject image.Image, out *os.File) (
	err error) {

	// set image quality to 100%
	opt := jpeg.Options{
		Quality: 100,
	}

	// write out the image object into the new PNG file
	err = jpeg.Encode(out, imageObject, &opt)
	if err != nil {
		return err
	}

	return nil
}

// exportPNG exporting image object to png file
// will return error status & file object
func exportPNG(imageObject image.Image, out *os.File) (
	err error) {

	// write out the image object into the new PNG file
	err = png.Encode(out, imageObject)
	if err != nil {
		return err
	}

	return nil
}

// exportGIF exporting image object to png file
// will return error status & file object
func exportGIF(imageObject image.Image, out *os.File) (
	err error) {

	// set image quality
	opt := gif.Options{
		NumColors: 256,
	}

	// write out the image object into the new GIF file
	err = gif.Encode(out, imageObject, &opt)
	if err != nil {
		return err
	}

	return nil
}
