package main

import (
	"fmt"
	goWatermark "github.com/michaelwp/go-watermark"
	"image/color"
)

func main() {
	err := goWatermark.AddWatermark(
		&goWatermark.Watermark{
			Image:      "input.jpg",
			OutputFile: "output.jpg",
			Text:       "GO WATERMARK",
			Position: goWatermark.Position{
				PosX:  0,
				PosY:  0,
				PosAY: 0,
				PosAX: 0,
			},
			Font: goWatermark.Font{
				FontSize: 20,
				FontName: "arial.ttf",
			},
			Color: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 70,
			},
			Align: goWatermark.AlignCenter,
			Repeat: goWatermark.Repeat{
				RepY:        1000,
				RepX:        15,
				WordSpacing: 0,
			},
			LineSpacing: 50,
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Watermark added successfully!")
	}
}
