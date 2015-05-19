package colormap

import (
	"sort"
)

func (cmap *ColorMap) SortedIndexes() []int {
	sortedIndexes := make([]int, len(cmap))
	for i := range cmap {
		sortedIndexes[i] = i
	}
	sorter := colorMapSorter{
		Cmap:          *cmap,
		SortedIndexes: sortedIndexes,
	}

	sort.Sort(sorter)
	return sorter.SortedIndexes
}

type colorMapSorter struct {
	Cmap          ColorMap
	SortedIndexes []int
}

func (s colorMapSorter) Len() int {
	return len(s.SortedIndexes)
}

func (s colorMapSorter) Swap(i, j int) {
	s.SortedIndexes[i], s.SortedIndexes[j] = s.SortedIndexes[j], s.SortedIndexes[i]
}

func (s colorMapSorter) Less(i, j int) bool {
	// We need to sort largest to smallest
	return s.Cmap.ColorTotalAt(s.SortedIndexes[j]) < s.Cmap.ColorTotalAt(s.SortedIndexes[i])
}
