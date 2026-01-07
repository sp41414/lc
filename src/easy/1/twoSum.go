package twoSum

func TwoSum(nums []int, target int) []int {
	indexes := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if j, found := indexes[diff]; found {
			return []int{j, i}
		}
		indexes[v] = i
	}
	return []int{}
}
