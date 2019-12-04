package main

type ListNode struct {
	Val  int
	next *ListNode
}

func merge(p1, p2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.next = p1
			p = p1
			p1 = p1.next
		} else {
			p.next = p2
			p = p2
			p2 = p2.next
		}
	}

	if p1 != nil {
		p.next = p1
	} else {
		p.next = p2
	}
	return head.next
}

func sortList(head *ListNode, n int) *ListNode {
	if n <= 1 {
		return head
	}
	p := head
	q := head
	for i := 0; i < n/2; i++ {
		q = p
		p = p.next
	}
	q.next = nil

	left := sortList(head, n/2)
	right := sortList(p, n-n/2)
	return merge(left, right)
}

func sort(head *ListNode) *ListNode {
	cnt := 0
	for p := head; p != nil; p = p.next {
		cnt++
	}

	return sortList(head, cnt)
}
