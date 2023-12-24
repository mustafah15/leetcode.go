package two_pointers

func maxArea(height []int) int {
	var mArea = 0
	var low = 0
	var high = len(height) - 1

	for low < high {
		area := min(height[low], height[high]) * (high - low)

		if area > mArea {
			mArea = area
		}
		if height[low] > height[high] {
			high--
		} else {
			low++
		}
	}

	return mArea
}
