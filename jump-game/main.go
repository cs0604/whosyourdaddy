package jump_game

//[3 2 1 0 4]
//[3 2 1 1 4]
//[2,0,0]

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//动态规划可以做，关键是状态转移方程需要设计好，这里以剩余跳力来做状态
func canJump(input []int) bool {
	var dp = make([]int, len(input)) //剩余跳力
	for i := 1; i < len(input); i++ {
		dp[i] = max(dp[i-1], input[i-1]) - 1
		if dp[i] < 0 {
			return false
		}
	}
	return true
}

//贪心算法，当前位置如果<=max，说明当前位置可达，然后判断是否需要更新max (与当前位置能跳跃的位置相比）
func canJump2(input []int) bool {
	if len(input) == 0 {
		return false
	}

	max := 0

	for i := 0; i < len(input); i++ {
		if i <= max {
			if max < i+input[i] {
				max = i + input[i]
			}
		}
		if max >= len(input)-1 {
			return true
		}
	}

	return false

}
