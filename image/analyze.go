package image

import (
	"image"
	"image/color"
)

func AvgRGBA(img image.Image) color.RGBA {
	bounds := img.Bounds()
	var _r, _g, _b, _a uint32

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			_a += (a >> 24 & 0xFF)
			_r += (r >> 16 & 0xFF)
			_g += (g >> 8 & 0xFF)
			_b += (b & 0xFF)
		}
	}

	pixels := bounds.Max.X * bounds.Max.Y

	return color.RGBA{
		uint8(int(_r) / pixels),
		uint8(int(_g) / pixels),
		uint8(int(_b) / pixels),
		uint8(int(_a) / pixels),
	}
}
