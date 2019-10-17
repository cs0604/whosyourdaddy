package Maximum_Depth_of_Binary_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	return visit(root, 0)
}

func visit(node *TreeNode, dep int) int {
	if node == nil {
		return dep
	}

	left := visit(node.Left, dep+1)
	right := visit(node.Right, dep+1)

	if left > right {
		return left
	}

	return right
}
