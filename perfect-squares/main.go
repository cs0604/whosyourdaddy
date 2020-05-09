package perfect_squares

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func numSquares(n int) int {

	var base []int
	i := 1
	for {
		s := i * i
		if s == n {
			return 1
		}
		if s > n {
			break
		}

		base = append(base, s)
		i++
	}

	var dp = make([]int, n+1)
	dp[0] = 0

	for i := 1; i < n+1; i++ {
		dp[i] = n + 1
		for j := 0; j < len(base); j++ {
			if i >= base[j] {
				dp[i] = min(dp[i], dp[i-base[j]]+1)
			}
		}
	}

	return dp[n]

}
