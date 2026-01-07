package removeDuplicatesSorted

func RemoveDuplicates(nums []int) int {
	k := 1
	pr := 0
	pw := 1

	for pr < len(nums) {
		if nums[pr] == nums[pw-1] {
			pr++
			continue
		}
		nums[pw] = nums[pr]
		pw++
		k++
	}

	return k
}
