package cie

import (
	"image/color"
	"math"

	mcolor "github.com/mrap/mosaic/color"
)

const (
	rgbCompandLine float64 = 0.04045
)

type XYZ struct {
	X float32
	Y float32
	Z float32
}

func XYZFromRGB(c color.RGBA) XYZ {
	cR, cG, cB := mcolor.NormalizedRGB(c)

	r := inverseSRGBCompand(cR)
	g := inverseSRGBCompand(cG)
	b := inverseSRGBCompand(cB)

	// Observer. = 2°, Illuminant = D65
	return XYZ{
		r*0.4124 + g*0.3576 + b*0.1805,
		r*0.2126 + g*0.7152 + b*0.0722,
		r*0.0193 + g*0.1192 + b*0.9505,
	}
}

func inverseSRGBCompand(v uint8) float32 {
	_v := float64(v)
	if _v > rgbCompandLine {
		_v = math.Pow((_v+0.055)/1.055, 2.4)
	} else {
		_v = _v / 12.92
	}

	return float32(_v * 100)
}
