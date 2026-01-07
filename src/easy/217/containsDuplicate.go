package containsDuplicate

func ContainsDuplicate(nums []int) bool {
	exists := make(map[int]struct{})

	for _, value := range nums {
		if _, ok := exists[value]; ok {
			return true
		} else {
			exists[value] = struct{}{}
		}
	}
	return false
}
