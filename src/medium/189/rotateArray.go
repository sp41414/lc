package rotateArray

import "slices"

func Rotate(nums []int, k int) {
	n := len(nums)

	k = k % n

	slices.Reverse(nums)

	for i := 0; i < k/2; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}

	for i := k; i < (n+k)/2; i++ {
		nums[i], nums[n+k-i-1] = nums[n+k-i-1], nums[i]
	}
}
