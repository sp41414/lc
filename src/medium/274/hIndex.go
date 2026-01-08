package hIndex

import "sort"

func HIndex(citations []int) int {
	hIndex := 0
	sort.Ints(citations)

	for i, v := range citations {
		if v >= len(citations[i:]) {
			hIndex = len(citations[i:])
			break
		}
	}
	return hIndex
}
