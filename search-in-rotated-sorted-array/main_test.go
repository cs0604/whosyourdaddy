package search_in_rotated_sorted_array

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				nums:   []int{1, 3, 5},
				target: 1,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums:   []int{3, 1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bsearch(t *testing.T) {
	type args struct {
		nums   []int
		low    int
		high   int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bsearch(tt.args.nums, tt.args.low, tt.args.high, tt.args.target); got != tt.want {
				t.Errorf("bsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
