package majorityElement

func MajorityElement(nums []int) int {
	c := 0
	element := 0
	for _, v := range nums {
		if c == 0 {
			element = v
			c++
		} else if v == element {
			c += 1
		} else {
			c -= 1
		}
	}

	c = 0
	for _, v := range nums {
		if v == element {
			c++
		}
		if c > len(nums)/2 {
			return element
		}
	}

	return 0
}
