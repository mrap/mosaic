package colormap_test

import (
	"image"
	"testing"

	. "github.com/mrap/mosaic/colormap"
	. "github.com/mrap/mosaic/image"
)

func imageColorMap(filepath string) (*ColorMap, image.Image) {
	img, err := ImageFromFile(filepath)
	if err != nil {
		return nil, nil
	}
	return NewColorMap(img), img
}

func TestColorsCount(t *testing.T) {
	cmap, img := imageColorMap("images/color_black.jpeg")
	if cmap.TotalColors != PixelsCount(img) {
		t.Error("ColorMap doesn't have correct number of TotalColors")
	}
}

func TestCreateFromImage(t *testing.T) {
	cmap, _ := imageColorMap("images/color_black.jpeg")
	if len(cmap.Get(0)) == 0 {
		t.Error("ColorMap didn't collect colors correctly from image.")
	}

	cmap, _ = imageColorMap("images/color_white.jpeg")
	if len(cmap.Get(255)) == 0 {
		t.Error("ColorMap didn't collect colors correctly from image.")
	}
}

func TestColorTotalAt(t *testing.T) {
	cmap, img := imageColorMap("images/color_black.jpeg")
	if cmap.ColorTotalAt(0) != PixelsCount(img) {
		t.Error("ColorTotalAt is not correct")
	}
}
