package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

//1 2 4
//1 3 4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var ret = &ListNode{}
	p := ret

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return ret.Next

}
