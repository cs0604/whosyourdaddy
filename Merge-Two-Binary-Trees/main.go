package Merge_Two_Binary_Trees

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}

	T := &TreeNode{}

	var l1, l2, r1, r2 *TreeNode

	if t1 != nil {
		l1 = t1.Left
		r1 = t1.Right
		T.Val += t1.Val
	}
	if t2 != nil {
		l2 = t2.Left
		r2 = t2.Right
		T.Val += t2.Val
	}

	T.Left = mergeTrees(l1, l2)
	T.Right = mergeTrees(r1, r2)

	return T
}
