package add_two_numbers

import (
	"reflect"
	"testing"
)

func buildList(input []int) *ListNode {
	head := &ListNode{}
	p := head

	for i := 0; i < len(input); i++ {
		p.Next = &ListNode{
			Val: input[i],
		}
		p = p.Next
	}
	return head.Next
}

func getListData(p *ListNode) []int {
	var res []int
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	return res
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				l1: buildList([]int{2, 4, 3}),
				l2: buildList([]int{5, 6, 4}),
			},
			want: []int{7, 0, 8},
		},
		{
			name: "2",
			args: args{
				l1: buildList([]int{8, 4, 3}),
				l2: buildList([]int{5, 6, 4}),
			},
			want: []int{3, 1, 8},
		},
		{
			name: "3",
			args: args{
				l1: buildList([]int{8, 4, 9, 9, 9}),
				l2: buildList([]int{5, 6}),
			},
			want: []int{3, 1, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(getListData(got), tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
