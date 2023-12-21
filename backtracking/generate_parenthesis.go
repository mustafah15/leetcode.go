package backtracking

func getParenthesis(output *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*output = append(*output, current)
		return
	}

	if open < max {
		getParenthesis(output, current+"(", open+1, close, max)
	}

	if open > close {
		getParenthesis(output, current+")", open, close+1, max)
	}
}

func generateParenthesis(n int) []string {
	var output = []string{}

	getParenthesis(&output, "", 0, 0, n)

	return output
}
