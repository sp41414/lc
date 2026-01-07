package sortArray

func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	firstHalf := nums[:mid]
	secondHalf := nums[mid:]

	firstSorted := SortArray(firstHalf)
	secondSorted := SortArray(secondHalf)

	merged := merge(firstSorted, secondSorted)

	return merged
}

func merge(firstHalf, secondHalf []int) []int {
	result := []int{}
	first, second := 0, 0

	for first < len(firstHalf) && second < len(secondHalf) {
		if firstHalf[first] <= secondHalf[second] {
			result = append(result, firstHalf[first])
			first++
		} else {
			result = append(result, secondHalf[second])
			second++
		}
	}

	if first < len(firstHalf) {
		result = append(result, firstHalf[first:]...)
	}

	if second < len(secondHalf) {
		result = append(result, secondHalf[second:]...)
	}

	return result
}
