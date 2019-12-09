package add_two_numbers

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	s := 0 //进位

	for l1 != nil && l2 != nil {
		c := (s + l1.Val + l2.Val) % 10
		s = (s + l1.Val + l2.Val) / 10
		p.Next = &ListNode{
			Val: c,
		}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	for s > 0 {
		if p.Next == nil {
			p.Next = &ListNode{
				Val: s,
			}
			break
		}
		p = p.Next
		c := (s + p.Val) % 10
		s = (s + p.Val) / 10
		p.Val = c
	}
	return head.Next
}
