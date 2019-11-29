package subsets

import (
	"reflect"
	"sort"
	"testing"
)

type R [][]int

func (r R) Len() int {
	return len([][]int(r))
}
func (r R) Less(i, j int) bool {
	if len(r[i]) != len(r[j]) {
		return len(r[i]) < len(r[j])
	}
	if len(r[i]) == 0 {
		return true
	}
	for k := 0; k < len(r[i]); k++ {
		if r[i][k] < r[j][k] {
			return true
		} else if r[i][k] > r[j][k] {
			return false
		}
	}
	return true
}

func (r R) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{3},
				{1},
				{2},
				{1, 2, 3},
				{1, 3},
				{2, 3},
				{1, 2},
				{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.args.nums)
			sort.Sort(R(got))
			sort.Sort(R(tt.want))
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
