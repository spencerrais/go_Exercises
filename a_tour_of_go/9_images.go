package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
	color         uint8
}

func (input Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, input.width, input.height)
}

func (input Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (input Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{150, 150, 1}
	pic.ShowImage(m)
}
