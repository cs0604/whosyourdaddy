package minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"ABCAAABEFD", "AD"},
			"ABEFD",
		},
		{
			"case2",
			args{"AD", "AD"},
			"AD",
		},
		{
			"case3",
			args{"ABCAAABEFD", "D"},
			"D",
		},
		{
			"case4",
			args{"ABCAAA", "ACB"},
			"ABC",
		},
		{
			"case5",
			args{"ABCAAABEFD", "AACBD"},
			"CAAABEFD",
		},
		{
			"case6",
			args{"ABCAAABFD", "AE"},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
