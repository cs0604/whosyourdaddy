package jump_game

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				input: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				input: []int{3, 2, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				input: []int{1, 1, 1, 0, 3},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				input: []int{0},
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				input: []int{2, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.input); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
