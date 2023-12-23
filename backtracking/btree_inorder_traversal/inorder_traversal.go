package backtracking

import backtracking "github.com/mustafah15/leetcode.go/backtracking"

func dfs(node *backtracking.TreeNode, visited *[]int) {
	if node != nil {
		if node.Left != nil {
			dfs(node.Left, visited)
		}
		*visited = append(*visited, node.Val)
		if node.Right != nil {
			dfs(node.Right, visited)
		}
	}
}

func inorder_traversal(root *backtracking.TreeNode) []int {
	var visited = []int{}
	dfs(root, &visited)
	return visited
}
