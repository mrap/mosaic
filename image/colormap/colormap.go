package colormap

import (
	"image"
	"image/color"

	libcolor "github.com/mrap/mosaic/color"
)

const (
	ColorMapBase int = 255
)

type ColorBase int

type ColorCounts map[color.RGBA]uint

type ColorMap struct {
	Bases       [ColorMapBase + 1]ColorCounts
	TotalColors uint
}

func (cmap *ColorMap) Add(c color.RGBA) {
	hexVal := libcolor.RGBHexValue(c)
	base := hexVal * uint32(ColorMapBase) / libcolor.MaxRGBValue
	if cmap.Bases[base] == nil {
		cmap.Bases[base] = make(ColorCounts)
	}
	cmap.Bases[base][c]++
	cmap.TotalColors++
}

func (cmap *ColorMap) Get(i ColorBase) ColorCounts {
	return cmap.Bases[i]
}

func NewColorMap(img image.Image) *ColorMap {
	cmap := new(ColorMap)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			cmap.Add(color.RGBA{
				A: (uint8)(a >> 8 & 0xFF),
				R: (uint8)(r >> 8 & 0xFF),
				G: (uint8)(g >> 8 & 0xFF),
				B: (uint8)(b >> 8 & 0xFF),
			})
		}
	}
	return cmap
}
