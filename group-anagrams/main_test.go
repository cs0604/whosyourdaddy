package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

type sar [][]string

func (s sar) Len() int {
	return len([][]string(s))
}
func (s sar) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
func (s sar) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func prepare(input [][]string) [][]string {

	for i := 0; i < len(input); i++ {
		sort.Strings(input[i])
	}

	sort.Sort(sar(input))

	return input
}

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				strs: []string{"eat", "aet", "tea", "abdd", "adbd"},
			},
			want: [][]string{{"eat", "aet", "tea"}, {"abdd", "adbd"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			t.Logf("%+v", got)
			t.Logf("%+v", tt.want)
			t.Logf("%+v", prepare(got))
			if !reflect.DeepEqual(prepare(got), prepare(tt.want)) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
