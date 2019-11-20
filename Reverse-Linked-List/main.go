package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	ret := reverse(head)
	head.Next = nil
	return ret
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	ret := reverse(p)
	p.Next = head
	return ret
}

func lprint(head *ListNode) {
	for q := head; q != nil; q = q.Next {
		fmt.Print(q.Val, "->")
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 0}
	p := head
	for i := 1; i < 5; i++ {

		p.Next = &ListNode{Val: i}
		p = p.Next

	}

	lprint(head)

	head = reverseList(head)

	lprint(head)

}
