package palindrome_partitioning

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		//{
		//	name: "1",
		//	args: args{
		//		s: "aab",
		//	},
		//	want: [][]string{
		//		{"aa","b"},
		//		{"a","a","b"},
		//	},
		//},
		{
			name: "2",
			args: args{
				s: "ababbbabbaba",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

//cbbbcc
//[["c","b","b","b","cc"],["c","b","b","b","c","c"],["c","b","bb","c"],["c","b","bb","c","c"],["c","bb","b","c"],["c","bb","b","c","c"],["c","bbb","cc"],["c","bbb","c","c"],["cbbbc","c"]]
//[["c","b","b","b","c","c"],["c","b","b","b","cc"],["c","b","bb","c","c"],["c","b","bb","cc"],["c","bb","b","c","c"],["c","bb","b","cc"],["c","bbb","c","c"],["c","bbb","cc"],["cbbbc","c"]]
