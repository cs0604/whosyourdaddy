package largest_number

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		//5 50 53 54 525  ->  50 525 52 53 54 5
		//50 525 524 53 54 5
		//3 30 34 341   ->  30 3 34 341  ->  30  3  341 34

		//20 1 -> 1 20 ->
		{
			name: "1",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "9534330",
		},
		{
			name: "2",
			args: args{
				nums: []int{20, 1},
			},
			want: "201",
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
