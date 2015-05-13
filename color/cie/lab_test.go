package cie_test

import (
	"image/color"
	"math"
	"testing"

	"github.com/mrap/mosaic/color/cie"
)

const epsilon = 0.0000001

func TestLabFromRGB(t *testing.T) {
	cRgb := color.RGBA{139, 164, 137, 1}
	cLab := cie.LabFromRGB(cRgb)

	if math.Abs(cLab.L-64.84017234415002) > epsilon {
		t.Error("Lab L is not correct")
	}

	if math.Abs(cLab.A-(-14.029491739167788)) > epsilon {
		t.Error("Lab A is not correct")
	}

	if math.Abs(cLab.B-11.14391132206023) > epsilon {
		t.Error("Lab B is not correct")
	}
}
