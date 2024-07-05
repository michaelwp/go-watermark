/**********************************************************
* 2024/07/05
* Author: Michael Putong
* This code free to use, share and modify
* Author not responsible for any damage caused by this code
***********************************************************/

package go_watermark

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"os"
	"strings"
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

type Position struct {
	PosX  float64
	PosY  float64
	PosAX float64
	PosAY float64
}

type Font struct {
	FontName string
	FontSize float64
}

type Watermark struct {
	Image      string
	OutputFile string
	Text       string
	Position
	Font
	Color       color.Color
	Align       Align
	LineSpacing float64
	Repeat
}

type Repeat struct {
	RepX, RepY, WordSpacing int
}

func AddWatermark(watermark *Watermark) error {
	bgImage, err := imageDecode(watermark.Image)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	fontByte := goregular.TTF
	if len(watermark.FontName) > 0 {
		fontByte, err = loadFont(watermark.FontName)
		if err != nil {
			return fmt.Errorf("error loading font %q: %v", watermark.FontName, err)
		}
	}

	font, err := truetype.Parse(fontByte)
	if err != nil {
		return fmt.Errorf("error in truetype.Parse: %v", err)
	}

	DrawWatermark(font, watermark, dc, float64(imgWidth), float64(imgHeight))

	err = dc.SavePNG(watermark.OutputFile)
	if err != nil {
		return fmt.Errorf("error saving image: %v", err)
	}

	return nil
}

func DrawWatermark(font *truetype.Font, watermark *Watermark, dc *gg.Context, imgWidth, imgHeight float64) {
	maxWidth := imgWidth - 60.0
	posY := int(watermark.PosY)
	if watermark.RepY < 2 {
		watermark.RepY = 0
	}

	if watermark.RepX < 1 {
		watermark.RepX = 1
	}

	y := float64(posY)

	for divY := 0; divY <= watermark.RepY-1; divY++ {
		wordSpaces := strings.Repeat(" ", watermark.WordSpacing)
		repTextX := strings.Repeat(watermark.Text+wordSpaces, watermark.RepX)

		face := truetype.NewFace(font, &truetype.Options{Size: watermark.FontSize})
		dc.SetFontFace(face)
		dc.SetColor(watermark.Color)
		dc.DrawStringWrapped(
			repTextX,
			watermark.PosX, y, watermark.PosAX, watermark.PosAY,
			maxWidth,
			watermark.LineSpacing,
			gg.Align(watermark.Align),
		)

		y += watermark.LineSpacing
	}
}

func imageDecode(imageFile string) (image.Image, error) {
	imgFile, err := os.Open(imageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}

	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image file: %v", err)
	}

	return img, nil
}

func loadFont(fontFile string) ([]byte, error) {
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %v", err)
	}

	return fontBytes, nil
}
