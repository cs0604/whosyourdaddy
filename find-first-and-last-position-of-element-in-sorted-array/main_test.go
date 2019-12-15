package find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 2, 2, 2, 3, 4},
				target: 2,
			},
			want: []int{1, 3},
		},
		{
			name: "2",
			args: args{
				nums:   []int{1, 2, 2, 2, 3, 4},
				target: 3,
			},
			want: []int{4, 4},
		},
		{
			name: "3",
			args: args{
				nums:   []int{1, 2, 2, 2, 3, 4},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "4",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "5",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
