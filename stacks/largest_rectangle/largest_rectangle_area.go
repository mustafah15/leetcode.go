package stacks

func largestRectangleArea(heights []int) int {
	var maxArea = 0
	// creating a stack to hold a pair of elements
	// index and hight of that index
	stack := [][2]int{}

	for i := range heights {
		// setting the start position of that area
		start := i
		// if the stack is not empty
		// and if the top item in the stack height is larger than hight  we just reached
		for len(stack) > 0 && stack[len(stack)-1][1] > heights[i] {
			lastIndex := stack[len(stack)-1][0]

			// trying to extend the hight that we are at backward to check if this the largest area
			// i - index is the width and the hight is the hight of the large height
			maxArea = max(maxArea, (i-stack[len(stack)-1][0])*stack[len(stack)-1][1])

			// since we know this hight is greater than the hight
			// we are visiting we can extend our area backward
			start = lastIndex
			// pop the height from the stack
			stack = stack[:len(stack)-1]
		}
		// adding another pair to the stack
		// from the start item
		stack = append(stack, [2]int{start, heights[i]})
	}

	// we also need to compute the area for items that left in the stack
	// so we calculate their area until the end of the histogram
	total := len(heights)
	for i := range stack {
		maxArea = max(maxArea, ((total - stack[i][0]) * stack[i][1]))
	}

	return maxArea
}
