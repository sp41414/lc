package medianSorted

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)

	n := len(merged)
	if n%2 == 0 {
		return float64(merged[n/2]+merged[n/2-1]) / 2.0
	} else {
		return float64(merged[n/2])
	}
}
