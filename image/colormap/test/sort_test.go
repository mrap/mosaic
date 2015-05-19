package colormap_test

import (
	"image/color"
	"testing"

	. "github.com/mrap/mosaic/image/colormap"
)

func TestSortedIndexes(t *testing.T) {
	cmap := new(ColorMap)
	whiteColor := color.RGBA{255, 255, 255, 0}
	cmap.Add(whiteColor)
	cmap.Add(whiteColor)
	blackColor := color.RGBA{0, 0, 0, 0}
	cmap.Add(blackColor)

	sortedIndexes := cmap.SortedIndexes()
	if sortedIndexes[0] != 255 {
		t.Error("SortedIndexes did not sort correctly. 255 should be first. Instead this was:", sortedIndexes[0])
	}

	if sortedIndexes[1] != 0 {
		t.Error("SortedIndexes did not sort correctly. 0 should be first. Instead this was:", sortedIndexes[1])
	}
}
