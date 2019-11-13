package best_time_to_buy_and_sell_stock

// 7 1 3 5 6 2
// 7 6 5 4 2 1
// 7 2 3 5 1 6
func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {

		if prices[i]-min > profit {
			profit = prices[i] - min
		}

		if prices[i] < min {
			min = prices[i]
		}

	}

	return profit

}
