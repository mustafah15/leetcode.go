package two_pointers

func max_area(height []int) int {
	var max_area = 0
	var low = 0
	var high = len(height) - 1

	for low < high {
		area := min(height[low], height[high]) * (high - low)

		if area > max_area {
			max_area = area
		}
		if height[low] > height[high] {
			high--
		} else {
			low++
		}
	}

	return max_area
}
