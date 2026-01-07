package removeDuplicatesSortedII

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	pr := 2
	pw := 2

	for pr < len(nums) {
		if nums[pr] != nums[pw-2] {
			nums[pw] = nums[pr]
			pw++
		}
		pr++
	}

	return pw
}
