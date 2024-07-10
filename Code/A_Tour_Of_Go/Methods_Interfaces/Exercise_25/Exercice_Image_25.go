package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct {
    x, y int
}

func (r *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, r.x, r.y)
}

func (r *Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
    m := &Image{256, 256}
    pic.ShowImage(m)
}