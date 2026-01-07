package shuffleArray

func Shuffle(nums []int, n int) []int {
	ans := []int{}
	p1 := 0
	p2 := n

	for p1 < n && p2 < len(nums) {
		ans = append(ans, nums[p1], nums[p2])
		p1++
		p2++
	}
	return ans
}
