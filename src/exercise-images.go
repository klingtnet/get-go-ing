package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rectangle{image.Point{0,0}, image.Point{512, 512}}
}

func (i Image) At(x, y int) color.Color {
    var r, g, b uint8 = uint8(x), uint8(x+y), uint8(y)
    if x > 256 && y > 256 {
        r,g,b = uint8(x+y), uint8(x), uint8(x+y)
    } else if x > 256 {
        r,g,b = uint8(x-y), uint8(x), uint8(y-x)
    } else if y > 256 {
        r,g,b = uint8(y-x), uint8(y), uint8(x-y)
    }
    return color.RGBA{r, g, b, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
