package permutations

import (
	"testing"
)

/**

123-> 123, 132
213-> 213,231
321->321,312


*/

func Test_permute(t *testing.T) {
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
				{1, 2, 3},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
			},
		},
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: [][]int{
				{1, 2, 3},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			t.Logf("len:%v", len(got))
			//if  !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("permute() = %v, want %v", got, tt.want)
			//}
		})
	}
}
