package find_all_duplicates_in_an_array

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 5, 6, 2, 3},
			},
			want: []int{2},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 6, 2, 3},
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
