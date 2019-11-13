package best_time_to_buy_and_sell_stock

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{7, 1, 2, 5, 4, 3, 6},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{7, 3, 4, 5, 1, 6},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
