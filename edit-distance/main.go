package edit_distance

func minDistance(word1 string, word2 string) int {

	var dp = make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {

			cost := 1
			if word1[i-1] == word2[j-1] {
				cost = 0
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)

		}
	}

	return dp[len(word1)][len(word2)]

}

func min(arg ...int) int {
	x := arg[0]
	for i := 1; i < len(arg); i++ {
		if x > arg[i] {
			x = arg[i]
		}
	}
	return x
}
