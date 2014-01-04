package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image interface {
    // ColorModel returns the Image's color model.
    ColorModel() color.Model
    // Bounds returns the domain for which At can return non-zero color.
    // The bounds do not necessarily contain the point (0, 0).
    Bounds() image.Rectangle
    // At returns the color of the pixel at (x, y).
    // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
    // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
    At(x, y int) color.Color
}

type MyImage struct {
    W int
    H int
}

func (m MyImage) ColorModel() color.Model {
    return color.RGBAModel
}

func (m MyImage) Bounds() image.Rectangle {
    return image.Rect(0, 0, m.W, m.H)
}

func (m MyImage) At(x, y int) color.Color {
    return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
    m := MyImage{10,10}
    pic.ShowImage(m)
}
