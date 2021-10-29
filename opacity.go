package go_watermark

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

// Opacity setup image transparency level
func Opacity(imageObject image.Image, alpha float64) (
	img image.Image, err error) {

	// validating image object
	if imageObject == nil {
		return nil, errors.New("image object required")
	}

	// validating alpha value, min = 0, max = 1
	if alpha < 0 {
		alpha = 0
	} else if alpha > 1 {
		alpha = 1
	}

	mapValues := func(value, start1, stop1, start2, stop2 float64) int {
		return int(start2 + (stop2-start2)*((value-start1)/(stop1-start1)))
	}

	bounds := imageObject.Bounds()
	mask := image.NewAlpha(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			r := mapValues(alpha, 1, 0, 0, 255)
			mask.SetAlpha(x, y, color.Alpha{A: uint8(255 - r)})
		}
	}

	maskImage := image.NewRGBA(bounds)
	draw.DrawMask(maskImage, bounds, imageObject, image.Pt(0, 0), mask, image.Pt(0, 0), draw.Over)
	return maskImage, nil
}
