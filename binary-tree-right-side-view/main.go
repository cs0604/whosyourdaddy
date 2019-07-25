package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	input := []string{"1", "2", "3", "null", "4", "5", "null"}

	root := buildTree(input)

	res := rightSideView(root)

	fmt.Printf("%+v", res)

}

func buildTree(input []string) *TreeNode {
	var arr []*TreeNode
	for _, str := range input {
		if str != "null" {
			val, _ := strconv.Atoi(str)
			arr = append(arr, &TreeNode{
				Val: val,
			})
		} else {
			arr = append(arr, nil)
		}
	}
	for i := 0; i < len(arr)/2; i++ {
		if i*2+1 < len(arr) {
			arr[i].Left = arr[i*2+1]
		}
		if i*2+2 < len(arr) {
			arr[i].Right = arr[i*2+2]
		}
	}
	return arr[0]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	type TreeNodeWrapper struct {
		node  *TreeNode
		level int
	}
	var queue = []*TreeNodeWrapper{{
		node:  root,
		level: 0,
	}}

	var lastLevel = 0
	var lastValue = root.Val

	var output []int

	//广度优先，按层遍历
	for {
		if len(queue) > 0 {
			curNode := queue[0]
			if curNode.node.Left != nil {
				queue = append(queue, &TreeNodeWrapper{
					node:  curNode.node.Left,
					level: curNode.level + 1,
				})
			}
			if curNode.node.Right != nil {
				queue = append(queue, &TreeNodeWrapper{
					node:  curNode.node.Right,
					level: curNode.level + 1,
				})
			}
			queue = queue[1:]

			if curNode.level > lastLevel {
				//输出上一层最后一个节点
				output = append(output, lastValue)
				lastLevel = curNode.level
			}
			lastValue = curNode.node.Val
		} else {
			break
		}
	}
	//输出最后一个节点
	output = append(output, lastValue)
	return output
}
