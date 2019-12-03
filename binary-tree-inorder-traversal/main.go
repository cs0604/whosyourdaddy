package binary_tree_inorder_traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

}

func inorder1(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder1(root.Left, res)
	*res = append(*res, root.Val)
	inorder1(root.Right, res)
}

func inorder2(root *TreeNode) []int {

	var stack []*TreeNode

	var res []int

	if root == nil {
		return nil
	}
	node := root
	for len(stack) > 0 || node != nil {
		if node == nil {
			node = stack[len(stack)-1]
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			node = node.Right
		} else {
			//进栈
			stack = append(stack, node)
			node = node.Left
		}

	}

	return res

}
