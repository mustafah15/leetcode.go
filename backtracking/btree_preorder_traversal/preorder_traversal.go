package backtracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, visited *[]int) {
	if node != nil {
		*visited = append(*visited, node.Val)
		if node.Left != nil {
			dfs(node.Left, visited)
		}
		if node.Right != nil {
			dfs(node.Right, visited)
		}
	}
}

func preorder_traversal(root *TreeNode) []int {
	var visited = []int{}
	dfs(root, &visited)

	return visited
}
