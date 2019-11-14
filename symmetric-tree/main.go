package symmetric_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkEq(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		if t1.Val == t2.Val {
			return checkEq(t1.Left, t2.Right) && checkEq(t1.Right, t2.Left)
		}
	}
	return false
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return checkEq(root.Left, root.Right)
}
