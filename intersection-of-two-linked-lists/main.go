package intersection_of_two_linked_lists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	var nodes = make(map[*ListNode]int)

	for p := headA; p != nil; p = p.Next {
		nodes[p] = 1
	}

	for p := headB; p != nil; p = p.Next {
		if _, ok := nodes[p]; ok {
			return p
		}
	}

	return nil
}
