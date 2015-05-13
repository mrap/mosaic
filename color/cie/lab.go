package cie

import (
	"image/color"
	"math"
)

// Observer= 2°, Illuminant= D65
var xyzRef = XYZ{
	X: 95.047,
	Y: 100.000,
	Z: 108.883,
}

type Lab struct {
	L int8
	A int8
	B int8
}

func LabFromRGB(c color.RGBA) Lab {
	return LabFromXYZ(XYZFromRGB(c))
}

func LabFromXYZ(c XYZ) Lab {
	x := convertLabFromXYZ(c.X / xyzRef.X)
	y := convertLabFromXYZ(c.Y / xyzRef.Y)
	z := convertLabFromXYZ(c.Z / xyzRef.Z)

	return Lab{
		L: int8(116*y - 16),
		A: int8(500 * (x - y)),
		B: int8(200 * (y - z)),
	}
}

const (
	epsilon = 0.008856
	kappa   = 903.3
)

func convertLabFromXYZ(v float32) float32 {
	if v > epsilon {
		return float32(math.Pow(float64(v), 1/3))
	} else {
		return kappa*v + 16/116
	}
}
