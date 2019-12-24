package partition_equal_subset_sum

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
