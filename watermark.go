package go_watermark

import (
	"errors"
	"image"
	"image/draw"
)

type Position int

const (
	Tiles Position = iota
	TopLeft
	TopRight
	TopMiddle
	CenterLeft
	CenterRight
	CenterMiddle
	BottomLeft
	BottomRight
	BottomMiddle
)

func (p Position) String() string {
	switch p {
	case TopLeft:
		return "TOP_LEFT"
	case TopRight:
		return "TOP_RIGHT"
	case TopMiddle:
		return "TOP_MIDDLE"
	case CenterLeft:
		return "CENTER_LEFT"
	case CenterRight:
		return "CENTER_RIGHT"
	case CenterMiddle:
		return "CENTER_MIDDLE"
	case BottomLeft:
		return "BOTTOM_LEFT"
	case BottomRight:
		return "BOTTOM_RIGHT"
	case BottomMiddle:
		return "BOTTOM_MIDDLE"
	default:
		return "TILES"
	}
}

type Option struct {
	Position Position
	Opacity  float64
}

// mergeImg uses go standard image.Image to get the watermark image and original image that want to watermark,
// the position of the watermark has to provide in image.Point then it'll return the watermarked image output
func mergeImg(watermark image.Image, original image.Image, position *image.Point) (
	img image.Image, err error) {

	var defaultPosition = image.Pt(0, 0)

	// validating image
	if watermark == nil || original == nil {
		msg := "watermark & original image required"
		return nil, errors.New(msg)
	}

	// validating position
	if position == nil {
		position = &defaultPosition
	}

	// setup bounds for watermark and original image
	// for watermark image bounds will add position coordinate
	originalImgBounds := original.Bounds()
	watermarkImgBounds := watermark.Bounds().Add(*position)

	// setup destination image
	destinationImgRGBA := image.NewRGBA(originalImgBounds)

	draw.Draw(destinationImgRGBA, originalImgBounds, original, image.Pt(0, 0), draw.Src)
	draw.Draw(destinationImgRGBA, watermarkImgBounds, watermark, image.Pt(0, 0), draw.Over)

	return destinationImgRGBA, nil
}

func Watermark(original string, watermark string, option *Option) {

}
