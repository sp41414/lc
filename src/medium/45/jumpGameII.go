package jumpGameII

func Jump(nums []int) int {
	curr_jump := 0
	curr_reach := 0
	jumps := 0

	for i, v := range nums[:len(nums)-1] {
		if i+v > curr_reach {
			curr_reach = i + v
		}
		if i == curr_jump {
			jumps++
			curr_jump = curr_reach
		}
	}

	return jumps
}
