package colormap

func (cmap *ColorMap) TopColorBases(minPercent float32) []ColorBase {
	if minPercent > 1 {
		minPercent /= 100
	}

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
