package colormap_test

import (
	"image/color"
	"testing"

	. "github.com/mrap/mosaic/colormap"
)

func TestTopColorBases(t *testing.T) {
	cmap := new(ColorMap)
	whiteColor := color.RGBA{255, 255, 255, 0}
	cmap.Add(whiteColor)
	cmap.Add(whiteColor)
	blackColor := color.RGBA{0, 0, 0, 0}
	cmap.Add(blackColor)

	bases := cmap.TopColorBases(66.6)
	if len(bases) > 1 {
		t.Error("TopColorBases returned too many bases")
	}
	if bases[0] != ColorBase(255) {
		t.Error("TopColorBases did not return correct base")
	}
}
