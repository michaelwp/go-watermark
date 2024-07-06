# go-watermark
This package allows you to add customizable watermarks to images using the Go programming language. 
It provides functionalities for positioning, repeating text, and adjusting font properties.

### Installation

```sh
go get github.com/michaelwp/go-watermark
```

### Example

```go
package main

import (
   "fmt"
   goWatermark "github.com/michaelwp/go-watermark"
   "image/color"
)


func main() {
    watermark := &Watermark{
        Image:      "input.png",
        OutputFile: "output.png",
        Text:       "Sample Watermark",
        Position: Position{
            PosX:  100,
            PosY:  100,
            PosAX: 0.5,
            PosAY: 0.5,
        },
        Font: Font{
            FontName: "path/to/font.ttf",
            FontSize: 36,
        },
        Color:       color.RGBA{255, 0, 0, 255},
        Align:       AlignCenter,
        LineSpacing: 1.5,
        Repeat: Repeat{
            RepX:        1,
            RepY:        1,
            WordSpacing: 2,
        },
    }

    if err := AddWatermark(watermark); err != nil {
        fmt.Println("Error adding watermark:", err)
    } else {
        fmt.Println("Watermark added successfully!")
    }
}
```

This example demonstrates how to configure and apply a watermark to an image using the `go_watermark` package. Adjust the parameters as needed to fit your specific use case.

for detail explanation visit: [How to Add a Watermark onto an Image Using Go](https://www.goblog.dev/articles/33)

