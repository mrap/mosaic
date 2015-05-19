package colormap

import (
	"image/color"
	"math"

	"github.com/mrap/mosaic/unit"
)

const (
	minColorBaseDist = (ColorMapBase + 1) / 16
)

func (cmap *ColorMap) TopColorBases(minPercent float32) []ColorBase {
	minPercent = unit.Percentage(minPercent)

	var topBases []ColorBase

	totalColors := float32(cmap.TotalColors)

OuterLoop:
	for _, b := range cmap.SortedBases() {
		if float32(cmap.ColorTotalAt(b))/totalColors < minPercent {
			break
		}

		// Ensure minimum distance rule to give a more definitive distribution
		for _, tb := range topBases {
			if int(math.Abs(float64(tb-b))) < minColorBaseDist {
				continue OuterLoop
			}
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
