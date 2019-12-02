package kth_largest_element_in_an_array

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 7},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{7, 3, 3, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, 0, len(tt.args.nums)-1)
			t.Logf("%+v", tt.args.nums)
		})
	}
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 4, 2, 5, 3, 7},
				k:    1,
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 4, 2, 5, 3, 7},
				k:    2,
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				nums: []int{7, 6, 5, 4, 3, 2, 1},
				k:    2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
