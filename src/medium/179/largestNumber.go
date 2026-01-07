package largestNumber

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})

	result := ""

	for _, v := range nums {
		result += strconv.Itoa(v)
	}

	if result[0] == '0' {
		return "0"
	}

	return result
}
