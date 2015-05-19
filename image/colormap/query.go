package colormap

import (
	"image/color"

	"github.com/mrap/mosaic/unit"
)

func (cmap *ColorMap) TopColorBases(minPercent float32) []ColorBase {
	minPercent = unit.Percentage(minPercent)

	var topBases []ColorBase

	var totalColors float32 = float32(cmap.TotalColors)

	for _, b := range cmap.SortedBases() {
		if float32(cmap.ColorTotalAt(b))/totalColors < minPercent {
			break
		}
		topBases = append(topBases, b)
	}
	return topBases
}

const topPaletteStartingPercent float32 = 30.0

func (cmap *ColorMap) TopPalette(minColors int) color.Palette {
	var (
		topBases  []ColorBase
		topColors color.Palette
		percent   = topPaletteStartingPercent
	)

	for len(topBases) < minColors {
		topBases = cmap.TopColorBases(percent)
		percent -= percent / 2
	}

	for _, b := range topBases {
		topColors = append(topColors, cmap.AvgColorAt(b))
	}

	return topColors
}
