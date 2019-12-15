package linked_list_cycle_ii

import (
	"reflect"
	"testing"
)

func createList(arr []int, pos int) (*ListNode, *ListNode) {
	head := &ListNode{}
	p := head
	var cycle *ListNode
	for i := 0; i < len(arr); i++ {
		p.Next = &ListNode{
			Val: arr[i],
		}
		p = p.Next
		if pos == i {
			cycle = p
		}
	}

	p.Next = cycle

	return head.Next, cycle
}

func Test_detectCycle(t *testing.T) {
	type tcase struct {
		name string
		args *ListNode
		want *ListNode
	}
	var tests []*tcase
	p1, p2 := createList([]int{3, 2, 0, 4}, 1)
	tests = append(tests, &tcase{
		name: "1",
		args: p1,
		want: p2,
	})
	p1, p2 = createList([]int{3, 2, 0, 4}, -1)
	tests = append(tests, &tcase{
		name: "2",
		args: p1,
		want: p2,
	})
	p1, p2 = createList([]int{1, 2}, 0)
	tests = append(tests, &tcase{
		name: "3",
		args: p1,
		want: p2,
	})
	p1, p2 = createList([]int{1}, 0)
	tests = append(tests, &tcase{
		name: "4",
		args: p1,
		want: p2,
	})
	p1, p2 = createList([]int{1}, -1)
	tests = append(tests, &tcase{
		name: "5",
		args: p1,
		want: p2,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
