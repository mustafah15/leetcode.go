package backtracking

import backtracking "github.com/mustafah15/leetcode.go/backtracking"

func dfs(node *backtracking.TreeNode) {
	if node == nil {
		return
	}

	left := node.Left
	right := node.Right

	node.Left = right
	node.Right = left

	if node.Left != nil {
		dfs(node.Left)
	}
	if node.Right != nil {
		dfs(node.Right)
	}
}

func invertTree(root *backtracking.TreeNode) *backtracking.TreeNode {
	dfs(root)
	return root
}
