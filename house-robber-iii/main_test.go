package house_robber_iii

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func createTree(str string) *TreeNode {
	var input []string
	input = strings.Split(str, ",")
	for i := 0; i < len(input); i++ {
		fmt.Printf("%s ", input[i])
	}
	fmt.Println()
	if len(input) > 0 {
		val, _ := strconv.Atoi(input[0])
		root := &TreeNode{
			Val: val,
		}
		build(root, 0, input)
		return root
	}

	return nil
}

func build(node *TreeNode, index int, input []string) {
	if node == nil {
		return
	}
	left := index*2 + 1
	right := index*2 + 2
	if left < len(input) && input[left] != "null" {
		val, _ := strconv.Atoi(input[left])
		tmp := &TreeNode{
			Val: val,
		}
		node.Left = tmp
	}
	if right < len(input) && input[right] != "null" {
		val, _ := strconv.Atoi(input[right])
		tmp := &TreeNode{
			Val: val,
		}
		node.Right = tmp
	}

	build(node.Left, left, input)
	build(node.Right, right, input)
}

func Test_rob(t *testing.T) {

	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: `3,2,3,null,3,null,1`,
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				input: `3,4,5,1,3,null,1`,
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.args.input)
			if got := rob(root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}

}
