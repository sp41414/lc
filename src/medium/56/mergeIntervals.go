package mergeIntervals

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort intervals by start time (least to greatest)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		// compare start time of current interval with last merged end time
		if intervals[i][0] <= merged[len(merged)-1][1] {
			// merge
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		} else {
			// no overlap, append
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
