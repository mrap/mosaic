package colormap

import (
	"image/color"
)

func (cmap *ColorMap) ColorTotalAt(i ColorBase) uint {
	return cmap.baseTotal[i]
}

func (cmap *ColorMap) AvgColorAt(i ColorBase) color.RGBA {
	var r, g, b, a uint

	for clr, count := range cmap.Get(i) {
		r += uint(clr.R) * count
		g += uint(clr.G) * count
		b += uint(clr.B) * count
		a += uint(clr.A) * count
	}

	total := float32(cmap.ColorTotalAt(i))
	return color.RGBA{
		uint8(float32(r) / total),
		uint8(float32(g) / total),
		uint8(float32(b) / total),
		uint8(float32(a) / total),
	}
}
