package diameter_of_binary_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxpath int

func diameterOfBinaryTree(root *TreeNode) int {
	maxpath = 0
	maxDepth(root)
	return maxpath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return -1
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	maxpath = max(maxpath, left+right+2)
	return max(left, right) + 1
}
