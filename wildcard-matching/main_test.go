package wildcard_matching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				s: "aa",
				p: "a",
			},
			false,
		},
		{
			"2",
			args{
				s: "aa",
				p: "*",
			},
			true,
		},
		{
			"3",
			args{
				s: "abcab",
				p: "a*a",
			},
			false,
		},
		{
			"4",
			args{
				s: "abcabc",
				p: "a*a?c",
			},
			true,
		},
		{
			"5",
			args{
				s: "cb",
				p: "?a",
			},
			false,
		},
		{
			"6",
			args{
				s: "adceb",
				p: "*a*b",
			},
			true,
		},
		{
			"7",
			args{
				s: "acdcb",
				p: "a*c?b",
			},
			false,
		},
		{
			"8",
			args{
				s: "ho",
				p: "ho****",
			},
			true,
		},
		{
			"9",
			args{
				s: "",
				p: "?",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
