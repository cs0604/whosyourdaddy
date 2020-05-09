package binary_tree_level_order_traversal

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	var queueNode []*TreeNode
	var result [][]int
	queueNode = append(queueNode, root)

	for len(queueNode) > 0 {
		levelSize := len(queueNode)

		var curVal []int
		for i := 0; i < levelSize; i++ {
			p := queueNode[i]
			curVal = append(curVal, p.Val)
			if p.Left != nil {
				queueNode = append(queueNode, p.Left)
			}
			if p.Right != nil {
				queueNode = append(queueNode, p.Right)

			}
		}

		result = append(result, curVal)

		queueNode = queueNode[levelSize:]
	}

	return result

}
