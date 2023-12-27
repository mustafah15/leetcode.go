package backtracking

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	dfs(0, target, candidates, &res, []int{})

	return res
}

func dfs(i int, rem int, candidates []int, res *[][]int, path []int) {
	l := len(candidates)

	// found a solution
	if rem == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	// reached the end of the array
	if i == l {
		return
	}

	if rem-candidates[i] >= 0 {
		// adding the current item to the path array
		// do
		path = append(path, candidates[i])
		// continue to find another candidates to meet the target with the current path
		// recurse
		dfs(i, rem-candidates[i], candidates, res, path)
		// remove/pop the candidates we have picked from the path so we can try another item
		// undo
		path = path[:len(path)-1]
	}

	// trying down the path with another candidate
	dfs(i+1, rem, candidates, res, path)
}
