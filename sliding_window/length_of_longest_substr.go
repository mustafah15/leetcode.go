package slidingwindow

func lengthOfLongestSubstring(s string) int {
	// defining a hash
	store := make(map[uint8]int)
	var left, right, result int

	for right < len(s) {
		var r = s[right]
		// adding charcter to the hash
		store[r] += 1

		// if it's duplicate
		for store[r] > 1 {
			l := s[left]
			// getting the left char
			// decrement it from the hash
			store[l] -= 1
			// moving the left one step forward
			left += 1
		}
		// making the result to hold the miximam
		// length found between the current result and the new window size
		result = max(result, right-left+1)
		// incrementing the right to take try one more char
		right += 1
	}

	return result
}
