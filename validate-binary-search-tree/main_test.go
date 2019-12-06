package validate_binary_search_tree

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

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root: createTree("1,1"),
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				root: createTree("1,2,3,null,4"),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				root: createTree("2,1,3,null,null,null,4"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
