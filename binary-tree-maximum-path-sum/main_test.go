package binary_tree_maximum_path_sum

import (
	"encoding/json"
	"fmt"
	"testing"
)

func buildTree(input string) *TreeNode {

	//bfs build tree
	return nil
}

func Test_unmarshal(t *testing.T) {
	input := "[1,2,3,null,0,2]"
	var data1 []int
	var data2 []string
	json.Unmarshal([]byte(input), &data1)
	json.Unmarshal([]byte(input), &data2)

	fmt.Printf("data1:%+v\n", data1)
	fmt.Printf("data2:%+v\n", data2)
}

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
