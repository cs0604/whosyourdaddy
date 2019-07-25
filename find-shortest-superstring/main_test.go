package find_shortest_superstring

import "testing"

func Test_shortestSuperstring(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"0",
			args{
				A: []string{"alex", "lexa"},
			},
			"alexa",
		},
		{
			"1",
			args{
				A: []string{"alex", "loves", "leetcode"},
			},
			"alexlovesleetcode",
		},
		{
			"2",
			args{
				A: []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"},
			},
			"gctaagttcatgcatc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSuperstring(tt.args.A); len(got) != len(tt.want) {
				t.Errorf("shortestSuperstring() = %v(%v), want %v(%v)", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}
