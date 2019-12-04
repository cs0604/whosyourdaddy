package palindrome_linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//O（1）空间需要逆置单链表的一部分再比较
func isPalindrome(head *ListNode) bool {

	count := 0

	for p := head; p.Next != nil; p = p.Next {
		count++
	}

	if count == 0 {
		return false
	}

	if count == 1 {
		return true
	}

	reverseCnt := count / 2

	var p, q, r *ListNode

	p = head
	q = p.Next
	r = q.Next

	//逆序单链表
	for i := 0; i < reverseCnt-1; i++ {
		q.Next = p
		p = q
		q = r
		r = r.Next
	}

	head.Next = nil

	//比较p1,p2两个链表的值是否相等
	p1 := p
	var p2 *ListNode
	if count%2 == 0 {
		p2 = q
	} else {
		p2 = r
	}

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return false

}
