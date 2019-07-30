package binary_tree_maximum_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]
func getMax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func maxPathSum(root *TreeNode) int {
	var maxVal = -2147483647
	var findMax func(node *TreeNode) int
	findMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getMax(0, findMax(node.Left))
		right := getMax(0, findMax(node.Right))

		maxVal = getMax(left+node.Val+right, maxVal)

		return getMax(left, right) + node.Val
	}

	findMax(root)

	return maxVal
}
