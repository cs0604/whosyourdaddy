package longest_consecutive_sequence

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			"case1", args{nums: []int{200, 3, 5, 1, 2, 1000}}, 3,
		},
		{
			"case2", args{nums: []int{200, 3, 5, 1, 2, 1000, 4}}, 5,
		},
		{
			"case3", args{nums: []int{200, 3, 5, 1, 2, 1000, 6, 8, 7}}, 4,
		},
		{
			"case4", args{nums: []int{200, 3, 5, 1, 2, 1000, 6, 8, 7, 4}}, 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
