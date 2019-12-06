package validate_binary_search_tree

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//[10,5,15,null,null,6,20]
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}

func traverseInorder(node *TreeNode, arr *[]int) bool {
	if node == nil {
		return true
	}

	if !traverseInorder(node.Left, arr) {
		return false
	}

	if len(*arr) > 0 && node.Val <= (*arr)[len(*arr)-1] {
		return false
	}

	*arr = append(*arr, node.Val)
	if !traverseInorder(node.Right, arr) {
		return false
	}
	return true
}

func helper(node *TreeNode, low, high int) bool {
	if node == nil {
		return true
	}

	if node.Val <= low || node.Val >= high {
		return false
	}

	return helper(node.Left, low, node.Val) && helper(node.Right, node.Val, high)
}
