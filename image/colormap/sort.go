package colormap

import (
	"sort"
)

func (cmap *ColorMap) SortedBases() []ColorBase {
	sortedBases := make([]ColorBase, len(cmap))
	for i := range cmap {
		sortedBases[i] = ColorBase(i)
	}
	sorter := colorMapSorter{
		Cmap:        *cmap,
		SortedBases: sortedBases,
	}

	sort.Sort(sorter)
	return sorter.SortedBases
}

type colorMapSorter struct {
	Cmap        ColorMap
	SortedBases []ColorBase
}

func (s colorMapSorter) Len() int {
	return len(s.SortedBases)
}

func (s colorMapSorter) Swap(i, j int) {
	s.SortedBases[i], s.SortedBases[j] = s.SortedBases[j], s.SortedBases[i]
}

func (s colorMapSorter) Less(i, j int) bool {
	// We need to sort largest to smallest
	return s.Cmap.ColorTotalAt(s.SortedBases[j]) < s.Cmap.ColorTotalAt(s.SortedBases[i])
}
