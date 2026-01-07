package removeElement

func RemoveElement(nums []int, val int) int {
	p1 := 0
	p2 := len(nums) - 1
	k := len(nums)

	for p1 <= p2 {
		if nums[p1] == val {
			k--
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p2--
		} else {
			p1++
		}
	}

	return k
}
