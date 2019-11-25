package house_robber_iii

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robTree(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	left := robTree(root.Left)
	right := robTree(root.Right)
	var res [2]int
	res[0] = max(left[0], left[1]) + max(right[0], right[1])
	res[1] = root.Val + left[0] + right[0]
	return res
}
