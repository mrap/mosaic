package color_test

import (
	"image/color"
	"testing"

	. "github.com/mrap/mosaic/color"
)

func TestHexValue(t *testing.T) {
	rgb := color.RGBA{0, 163, 136, 0}
	if HexValue(rgb) != 41864 {
		t.Error("HexValue returns incorrect rgb hex value")
	}

}
