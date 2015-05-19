package mosaic_test

import (
	"testing"

	. "github.com/mrap/mosaic"
)

func TestNewMosaic(t *testing.T) {
	var (
		baseImagePath = "../colormap/test/images/color_black.jpeg"
		imgPaths      []string
		imgPathsCount = 4
	)

	for i := 0; i < imgPathsCount; i++ {
		imgPaths = append(imgPaths, "../colormap/test/images/color_white.jpeg")
	}

	m, err := NewMosaic(baseImagePath, imgPaths)
	if err != nil {
		t.Error(err)
	}

	if m.BaseImage == nil {
		t.Error("NewMosaic base image shouldn't be nil")
	}

	if len(m.Images) != imgPathsCount {
		t.Errorf("NewMosaic images should have %d images", imgPathsCount)
	}
}
