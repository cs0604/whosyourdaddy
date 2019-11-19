package path_sum_iii

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func createTree(str string) *TreeNode {
	var input []string
	json.Unmarshal([]byte(str), &input)
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

func Test_pathSum(t *testing.T) {
	type args struct {
		input string
		sum   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: `["10","5","-3","3","2","null","11","3","-2","null","1"]`,
				sum:   8,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				input: `["5","4","8","11","null","13","4","7","2","null","null","5","1"]`,
				sum:   22,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.args.input)
			if got := pathSum(root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
