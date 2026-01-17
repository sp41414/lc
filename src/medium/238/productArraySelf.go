package productArraySelf

import "slices"

func ProductExceptSelf(nums []int) []int {
	ans := slices.Repeat([]int{1}, len(nums))
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	right := nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		ans[i] *= right
		right *= nums[i]
	}

	return ans
}
