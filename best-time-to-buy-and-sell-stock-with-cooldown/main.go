package best_time_to_buy_and_sell_stock_with_cooldown

var min = -2147483648

func fillMin(dp []int) {
	for i := 0; i < len(dp); i++ {
		dp[i] = min
	}
}

func copyDp(from, to []int) {
	for i := 0; i < len(from); i++ {
		to[i] = from[i]
	}
}

func maxProfit(prices []int) int {
	//0 buy, 1 sell, 2 cooldown_hold 3 cooldown_unhold
	var dp1 = make([]int, 4)
	var dp2 = make([]int, 4)
	//初始状态
	fillMin(dp1)
	fillMin(dp2)
	dp1[3] = 0

	for i := 0; i < len(prices); i++ {
		//推导出新的状态，存放到dp2
		if dp1[0] > min {
			//buy ->  sell , cooldown_hold
			profit := dp1[0] + prices[i]
			if dp2[1] < profit {
				//update sell
				dp2[1] = profit
			}
			if dp2[2] < dp1[0] {
				//update cooldown_hold
				dp2[2] = dp1[0]
			}
		}

		if dp1[1] > min {
			//sell -> cooldown_unhold
			if dp2[3] < dp1[1] {
				dp2[3] = dp1[1]
			}
		}

		if dp1[2] > min {
			//cooldown_hold -> cooldown_hold, sell
			if dp2[2] < dp1[2] {
				dp2[2] = dp1[2]
			}
			profit := dp1[2] + prices[i]
			if dp2[1] < profit {
				dp2[1] = profit
			}
		}

		if dp1[3] > min {
			//cooldown_unhold -> cooldown_unhold, buy
			if dp2[3] < dp1[3] {
				dp2[3] = dp1[3]
			}
			profit := dp1[3] - prices[i]
			if dp2[0] < profit {
				dp2[0] = profit
			}
		}

		//更新dp1，dp2
		copyDp(dp2, dp1)
		fillMin(dp2)
	}

	max := 0
	for i := 0; i < len(dp1); i++ {
		if dp1[i] > max {
			max = dp1[i]
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit2(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s2 := make([]int, len(prices))

	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = -1 << 31

	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	return max(s0[len(prices)-1], s2[len(prices)-1])

}
