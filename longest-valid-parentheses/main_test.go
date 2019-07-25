package longest_valid_parentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				"(())",
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				"((())",
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				"())(())",
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				"())(",
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				"((()()))",
			},
			want: 8,
		},
		{
			name: "6",
			args: args{
				"()(()",
			},
			want: 2,
		},
		{
			name: "7",
			args: args{
				"()()",
			},
			want: 4,
		},

		{
			name: "8",
			args: args{
				")()(((())))(",
			},
			want: 10,
		},
		{
			name: "8",
			args: args{
				"(()(((()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
