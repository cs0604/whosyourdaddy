package symmetric_tree

import "testing"

func createTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	build(root, 0, nums)

	return root
}

func build(t *TreeNode, index int, nums []int) {
	if t == nil {
		return
	}

	leftIndex := index*2 + 1
	rightIndex := leftIndex + 1

	if leftIndex < len(nums) && nums[leftIndex] != -1 {
		tmp := &TreeNode{
			Val: nums[leftIndex],
		}
		t.Left = tmp
	}
	if rightIndex < len(nums) && nums[rightIndex] != -1 {
		tmp := &TreeNode{
			Val: nums[rightIndex],
		}
		t.Right = tmp
	}

	build(t.Left, leftIndex, nums)
	build(t.Right, rightIndex, nums)
}

func Test_isSymmetric(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "1",
			args: []int{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			name: "2",
			args: []int{1, 2, 2, -1, 3, -1, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.args)
			if got := isSymmetric(root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
