package minimum_path_sum

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {

	max := 2147483647

	M := len(grid)
	N := len(grid[0])

	dp := make([][]int, M+1)
	for i := 0; i < M+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i < M+1; i++ {
		dp[i][0] = max
	}
	for i := 1; i < N+1; i++ {
		dp[0][i] = max
	}

	for i := 1; i < M+1; i++ {
		for j := 1; j < N+1; j++ {
			if i == 1 && j == 1 {
				dp[i][j] = grid[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
			}
		}
	}

	return dp[M][N]

}
