package colormap_test

import (
	"image/color"
	"testing"

	. "github.com/mrap/mosaic/image/colormap"
)

func TestSortedBases(t *testing.T) {
	cmap := new(ColorMap)
	whiteColor := color.RGBA{255, 255, 255, 0}
	cmap.Add(whiteColor)
	cmap.Add(whiteColor)
	blackColor := color.RGBA{0, 0, 0, 0}
	cmap.Add(blackColor)

	sortedBases := cmap.SortedBases()
	if sortedBases[0] != 255 {
		t.Error("SortedBases did not sort correctly. 255 should be first. Instead this was:", sortedBases[0])
	}

	if sortedBases[1] != 0 {
		t.Error("SortedBases did not sort correctly. 0 should be first. Instead this was:", sortedBases[1])
	}
}
