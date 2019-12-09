package coin_change

func coinChange2(coins []int, amount int) int {

	var dp = make([]int, amount+1)

	helper(amount, coins, dp)

	return dp[amount]

}

func helper(amount int, coins []int, dp []int) int {

	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if dp[amount] != 0 {
		return dp[amount]
	}

	min := 1 << 31
	for i := 0; i < len(coins); i++ {
		s := helper(amount-coins[i], coins, dp)
		if s != -1 && min > s {
			min = s
		}
	}
	if min != 1<<31 {
		dp[amount] = min + 1
		return min + 1
	}
	dp[amount] = -1
	return -1
}

//1 2 5    11
//1 1 2 2 1 ... 11
func coinChange3(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	var dp = make([]int, amount+1)

	for i := 0; i < len(coins); i++ {
		if coins[i] < amount+1 {
			dp[coins[i]] = 1
		}
		for j := coins[i] + 1; j < amount+1; j++ {
			if dp[j-coins[i]] > 0 {
				if dp[j] == 0 || dp[j] > dp[j-coins[i]]+1 {
					dp[j] = dp[j-coins[i]] + 1
				}
			}
		}
	}

	if dp[amount] > 0 {
		return dp[amount]
	}

	return -1

}

func coinChange(coins []int, amount int) int {

	var dp = make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				if dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]

}
