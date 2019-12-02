package longest_increasing_subsequence

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 2, 5, 1, 6},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 3, 5, 6, 4, 9},
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
