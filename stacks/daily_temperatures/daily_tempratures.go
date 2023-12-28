package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	// stack to hold the last temperature weve seen that hasnt seen one bigger yet
	stack := []int{}

	for index, temperature := range temperatures {
		// if the stack is not empty and the current temperature larger
		// than the tmp in the last item in the stack
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			// last stack index
			lastIndex := stack[len(stack)-1]
			// get the difference between the last index
			// in the stack and the current one
			distance := index - lastIndex
			// put the result in the of the last index in result array
			result[lastIndex] = distance
			// now pop the last item of the stack
			stack = stack[:len(stack)-1]
		}

		// add the current index to the stack
		stack = append(stack, index)
	}

	return result
}
