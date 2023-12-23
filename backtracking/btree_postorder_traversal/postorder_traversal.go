package backtracking

import backtracking "github.com/mustafah15/leetcode.go/backtracking"

func dfs(node *backtracking.TreeNode, visited *[]int) {
	if node != nil {
		if node.Left != nil {
			dfs(node.Left, visited)
		}
		if node.Right != nil {
			dfs(node.Right, visited)
		}

		*visited = append(*visited, node.Val)
	}
}

func postorderTraversal(root *backtracking.TreeNode) []int {
	var visited = []int{}
	dfs(root, &visited)
	return visited
}
