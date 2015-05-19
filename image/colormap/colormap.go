package colormap

import (
	"image"
	"image/color"

	libcolor "github.com/mrap/mosaic/color"
)

const ColorMapBase int = 256

type ColorCounts map[color.RGBA]uint

type ColorMap [ColorMapBase]ColorCounts

func (cmap *ColorMap) Add(c color.RGBA) {
	hexVal := libcolor.HexValue(c)
	base := hexVal % uint32(ColorMapBase)
	if cmap[base] == nil {
		cmap[base] = make(ColorCounts)
	}
	cmap[base][c]++
}

func (cmap *ColorMap) Get(i int) ColorCounts {
	return cmap[i]
}

func NewColorMap(img image.Image) *ColorMap {
	cmap := new(ColorMap)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			cmap.Add(color.RGBA{
				A: (uint8)(a >> 24 & 0xFF),
				R: (uint8)(r >> 16 & 0xFF),
				G: (uint8)(g >> 8 & 0xFF),
				B: (uint8)(b & 0xFF),
			})
		}
	}
	return cmap
}
