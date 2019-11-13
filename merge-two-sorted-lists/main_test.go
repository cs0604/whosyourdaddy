package merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

func printList(node *ListNode) {
	for ; node != nil; node = node.Next {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func buildList(arr []int) *ListNode {
	var head = &ListNode{}
	p := head
	for i := 0; i < len(arr); i++ {
		p.Next = &ListNode{
			Val: arr[i],
		}
		p = p.Next
	}
	return head.Next
}

func Test_mergeTwoLists(t *testing.T) {

	l1 := buildList([]int{1, 2, 4})
	l2 := buildList([]int{1, 3, 4})

	p := mergeTwoLists(l1, l2)

	printList(p)

}
