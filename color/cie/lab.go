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
	L float64
	A float64
	B float64
}

func LabFromRGB(c color.RGBA) Lab {
	return LabFromXYZ(XYZFromRGB(c))
}

func LabFromXYZ(c XYZ) Lab {
	x := convertLabFromXYZ(c.X / xyzRef.X)
	y := convertLabFromXYZ(c.Y / xyzRef.Y)
	z := convertLabFromXYZ(c.Z / xyzRef.Z)

	return Lab{
		L: (116*y - 16),
		A: (500 * (x - y)),
		B: (200 * (y - z)),
	}
}

const (
	epsilon float64 = 0.008856
	kappa   float64 = 903.3
)

func convertLabFromXYZ(v float64) float64 {
	if v > epsilon {
		return math.Pow(v, 1.0/3.0)
	} else {
		return kappa*v + 16.0/116.0
	}
}
