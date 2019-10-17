package quicksort

import (
	"sort"
)

func Sort(data sort.Interface) {
	quicksort(data, 0, data.Len()-1)
}

func quicksort(data sort.Interface, lo int, hi int) {
	if lo < hi {
		pl, pr := partitioning(data, lo, hi)
		quicksort(data, lo, pl-1)
		quicksort(data, pr+1, hi)
	}
}

func partitioning(data sort.Interface, lo int, hi int) (int, int) {
	var pr int = lo
	var pl int = lo

	for i := lo + 1; i <= hi; i++ {
		if !data.Less(i, pr) && !data.Less(pr, i) {
			data.Swap(pr+1, i)
			pr++
		} else if data.Less(i, pr) {
			data.Swap(i, pl)
			data.Swap(i, pr+1)
			pl++
			pr++

		}
	}

	return pl, pr
}
