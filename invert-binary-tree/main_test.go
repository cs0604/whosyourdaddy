package invert_binary_tree

import (
	"fmt"
	"testing"
)

func checkTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return checkTree(t1.Left, t2.Right) && checkTree(t1.Right, t2.Left)
}

func printTreeByLevel(node *TreeNode) {

	if node == nil {
		return
	}
	var nodeArray []*TreeNode
	nodeArray = append(nodeArray, node)

	var i int
	for i = 0; i < len(nodeArray); i++ {
		cur := nodeArray[i]
		fmt.Printf("%v,", cur.Val)
		if cur.Left != nil {
			nodeArray = append(nodeArray, cur.Left)
		}
		if cur.Right != nil {
			nodeArray = append(nodeArray, cur.Right)
		}
	}
	fmt.Println()
}

func createTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	build(root, 0, nums)

	return root
}

func build(t *TreeNode, index int, nums []int) {
	if t == nil {
		return
	}

	leftIndex := index*2 + 1
	rightIndex := leftIndex + 1

	if leftIndex < len(nums) && nums[leftIndex] != -1 {
		tmp := &TreeNode{
			Val: nums[leftIndex],
		}
		t.Left = tmp
	}
	if rightIndex < len(nums) && nums[rightIndex] != -1 {
		tmp := &TreeNode{
			Val: nums[rightIndex],
		}
		t.Right = tmp
	}

	build(t.Left, leftIndex, nums)
	build(t.Right, rightIndex, nums)
}

func Test_invertTree(t *testing.T) {

	//[4,2,7,1,3,6,9]

	nums := []int{4, 2, 7, 1, 3, 6, 9, 10}

	root := createTree(nums)

	printTreeByLevel(root)

	p := invertTree(root)

	printTreeByLevel(p)

	r := invertTree(p)

	printTreeByLevel(r)

	if checkTree(root, p) {
		t.Fail()
	}
}
