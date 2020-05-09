package remove_nth_node_from_end_of_list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	p1 := head
	p2 := head

	pre := p2

	for i := 0; i < n-1; i++ {
		p1 = p1.Next
	}

	for p1.Next != nil {
		p1 = p1.Next
		pre = p2
		p2 = p2.Next
	}

	if pre == p2 {
		return pre.Next
	}

	pre.Next = p2.Next

	return head

}
