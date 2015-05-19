package colormap_test

import (
	"testing"
)

const waterdropImgPath = "images/waterdrop_500x500.jpg"

func BenchmarkImageFileToColorMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imageColorMap(waterdropImgPath)
	}
}

func BenchmarkImageFileToTopPalette(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmap, _ := imageColorMap(waterdropImgPath)
		cmap.TopPalette(3)
	}
}
