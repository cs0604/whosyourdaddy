package sort_list

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	return mergeSort(head)

}

func merge(p1, p2 *ListNode) *ListNode {
	head := &ListNode{}

	p := head

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	} else {
		p.Next = p2
	}

	return head.Next

}

func mergeSort(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head
	for p2 != nil {
		if p2.Next == nil || p2.Next.Next == nil {
			break
		}
		p2 = p2.Next.Next
		p1 = p1.Next
	}

	p2 = p1.Next
	p1.Next = nil

	p1 = mergeSort(head)
	p2 = mergeSort(p2)

	return merge(p1, p2)

}
