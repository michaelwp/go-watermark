package main

import (
	"fmt"
	"image/color"

	goWatermark "github.com/michaelwp/go-watermark"
)

func main() {
	err := goWatermark.AddWatermark(
		&goWatermark.Watermark{
			Image:      "input1.jpeg",
			OutputFile: "output.jpeg",
			Text:       "79995782-PTGLOBALPRADANASEJAHTERA-227",
			Position: goWatermark.Position{
				PosX:  0,
				PosY:  -50,
				PosAY: 10,
				PosAX: 0,
			},
			Font: goWatermark.Font{
				FontSize: 12,
			},
			Color: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 80,
			},
			Align: goWatermark.AlignCenter,
			Repeat: goWatermark.Repeat{
				RepY:        20,
				RepX:        10,
				WordSpacing: 0,
			},
			LineSpacing: 25,
			Rotate:      -30,
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Watermark added successfully!")
	}
}
