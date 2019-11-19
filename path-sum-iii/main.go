package path_sum_iii

//搞个map记住从root到当前节点的和cursum，每次查找的时候，只需要看cursum - targetsum的差值是否再map中

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	pre := make(map[int]int)
	pre[0] = 1
	return travel(pre, 0, sum, root)
}

func travel(pre map[int]int, cursum int, sum int, node *TreeNode) int {
	if node == nil {
		return 0
	}

	cursum += node.Val

	var count = 0

	if val, ok := pre[cursum-sum]; ok {
		count = val
	}

	if val, ok := pre[cursum]; ok {
		pre[cursum] = val + 1
	} else {
		pre[cursum] = 1
	}
	count = count + travel(pre, cursum, sum, node.Left) + travel(pre, cursum, sum, node.Right)

	pre[cursum] = pre[cursum] - 1

	return count
}
