package sort_list

import (
	"fmt"
	"testing"
)

func createList(arr []int) *ListNode {
	head := &ListNode{
		Val: arr[0],
	}
	p := head
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{
			Val: arr[i],
		}
		p = p.Next
	}
	return head
}

func printList(p *ListNode) {
	for ; p != nil; p = p.Next {
		fmt.Print(p.Val, "->")
	}
	fmt.Println()
}

func checkEqual(p1 *ListNode, val []int) bool {
	printList(p1)
	if p1 == nil && len(val) == 0 {
		return true
	}
	if p1 == nil || len(val) == 0 {
		return false
	}

	for i := 0; i < len(val); i++ {
		if p1.Val != val[i] {
			return false
		}
		p1 = p1.Next
	}

	return true

}

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				head: createList([]int{1, 3, 2, 4, 6, 5}),
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !checkEqual(got, tt.want) {
				t.Errorf("sortList() = %+v, want %v", got, tt.want)
			}
		})
	}
}
