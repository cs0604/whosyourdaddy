package maximal_square

func min(args ...int) int {
	v := args[0]
	for i := 1; i < len(args); i++ {
		if v > args[i] {
			v = args[i]
		}
	}
	return v
}

func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	res := 0
	pre := 0
	col := len(matrix[0])
	dp := make([]int, col)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				pre = dp[j]
				dp[j] = int(matrix[i][j] - '0') // 0 or 1
			} else {
				cur := min(dp[j], dp[j-1], pre) + 1
				pre = dp[j]
				dp[j] = cur
				if cur > res {
					res = cur
				}
			}
			if res < dp[j] {
				res = dp[j]
			}
		}
	}

	return res * res
}
