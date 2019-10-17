package quicksort

import (
	"sort"
	"testing"
)

func checkAndOutputError(t *testing.T, expectedVal interface{}, givenVal interface{}) {
	if expectedVal != givenVal {
		t.Errorf("Expected %d but was %d", expectedVal, givenVal)
	}
}

func TestSort(t *testing.T) {
	a := []int{5, 4, 2, 3, 1}
	Sort(sort.IntSlice(a))
	checkAndOutputError(t, 1, a[0])
	checkAndOutputError(t, 2, a[1])
	checkAndOutputError(t, 3, a[2])
	checkAndOutputError(t, 4, a[3])
	checkAndOutputError(t, 5, a[4])
}

func TestEmptyArr(t *testing.T) {
	a := []int{}
	Sort(sort.IntSlice(a))
	checkAndOutputError(t, 0, len(a))
}
