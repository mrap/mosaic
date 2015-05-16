package colormap_test

import (
	"testing"

	"github.com/mrap/mosaic/image"
	. "github.com/mrap/mosaic/image/colormap"
)

func imageColorMap(filepath string) *ColorMap {
	img, err := image.ImageFromFile(filepath)
	if err != nil {
		return nil
	}
	return NewColorMap(img)
}

func TestCreateFromImage(t *testing.T) {
	cmap := imageColorMap("images/color_black.jpeg")
	if len(cmap.Get(0)) == 0 {
		t.Error("ColorMap didn't collect colors correctly from image.")
	}

	cmap = imageColorMap("images/color_white.jpeg")
	if len(cmap.Get(255)) == 0 {
		t.Error("ColorMap didn't collect colors correctly from image.")
	}
}
