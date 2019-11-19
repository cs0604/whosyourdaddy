package path_sum_iv

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: []int{113, 215, 221},
			},
			want: 12,
		},
		{
			name: "2",
			args: args{
				input: []int{113, 221},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				input: []int{113, 215, 221, 314, 325, 333, 347},
			},
			want: 43,
		},
		{
			name: "4",
			args: args{
				input: []int{113, 215, 221, 314, 349},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.input); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
