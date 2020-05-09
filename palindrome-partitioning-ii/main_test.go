package palindrome_partitioning_ii

import "testing"

func Test_minCut(t *testing.T) {
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
				s: "abcbcd",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				s: "abcd",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				s: "cabcbaca",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
