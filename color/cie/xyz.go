package cie

import (
	"image/color"
	"math"
)

const (
	rgbCompandLine float64 = 0.04045
)

type XYZ struct {
	X float64
	Y float64
	Z float64
}

func XYZFromRGB(c color.RGBA) XYZ {
	r := inverseSRGBCompand(float64(c.R) / 255)
	g := inverseSRGBCompand(float64(c.G) / 255)
	b := inverseSRGBCompand(float64(c.B) / 255)

	// Observer. = 2°, Illuminant = D65
	return XYZ{
		r*0.4124 + g*0.3576 + b*0.1805,
		r*0.2126 + g*0.7152 + b*0.0722,
		r*0.0193 + g*0.1192 + b*0.9505,
	}
}

func inverseSRGBCompand(v float64) float64 {
	if v > rgbCompandLine {
		return math.Pow((v+0.055)/1.055, 2.4) * 100
	} else {
		return v / 12.92 * 100
	}
}
