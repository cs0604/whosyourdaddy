package palindrome_linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//O（1）空间需要逆置单链表的一部分再比较
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head

	pre := slow

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}

	//把第一个截断
	pre.Next = nil
	second := slow
	if fast != nil {
		second = slow.Next
	}

	//reverse second
	second = reverse(second)

	//compare

	p1 := head
	p2 := second

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1 == nil && p2 == nil

}

func reverse(head *ListNode) *ListNode {
	var pre, next *ListNode
	current := head

	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}
