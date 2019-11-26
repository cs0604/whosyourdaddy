package unique_paths

func uniquePaths(m int, n int) int {

	var dp [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[n-1][m-1]

}

func uniquePaths2(m int, n int) int {

	var dp1 = make([]int, m)
	var dp2 = make([]int, m)

	for i := 0; i < m; i++ {
		dp1[i] = 1
		dp2[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp2[j] = dp1[j] + dp2[j-1]
		}
		//swap
		dp1, dp2 = dp2, dp1
	}

	return dp1[m-1]

}

func uniquePaths3(m int, n int) int {

	var dp = make([]int, m)

	for i := 0; i < m; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[m-1]

}
