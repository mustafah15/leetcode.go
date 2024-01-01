package slidingwindow

func checkInclusion(s1 string, s2 string) bool {
	// checking the edge case that s1 is larger than s2
	if len(s1) > len(s2) {
		return false
	}

	l, matches := 0, [26]int{}
	for i := range s1 {
		// marking s1 chars as ones
		matches[s1[i]-97]++
	}

	for r := range s2 {
		// decrementing the chars of s2 from the matches
		matches[s2[r]-97]--

		if matches == [26]int{} {
			return true
		}

		// while the right pointer is larger or equal length of s1
		// mark the matches of left pointer on s2 +1
		// and increment the left one more time
		if r+1 >= len(s1) {
			matches[s2[l]-97]++
			l++
		}
	}

	return false
}
