package two_pointers

func two_sum(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1

	for low < high {
		if numbers[low]+numbers[high] > target {
			high--
		} else if numbers[low]+numbers[high] < target {
			low++
		} else {
			return []int{low + 1, high + 1}
		}
	}

	return []int{0, 0}
}
