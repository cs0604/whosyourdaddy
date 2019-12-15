package linked_list_cycle_ii

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	p1 := head
	p2 := head
	for {
		if p1 == nil {
			return nil
		}
		if p2 == nil || p2.Next == nil {
			return nil
		}
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			break
		}
	}

	if p1 == nil {
		return nil
	}

	p2 = head

	for {
		if p2 == p1 {
			return p2
		}

		p2 = p2.Next
		p1 = p1.Next
	}

}
