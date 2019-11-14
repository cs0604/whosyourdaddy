package climbing_stairs

func climbStairs(n int) int {
	var fn = make([]int, n+1)
	fn[0] = 1

	for i := 1; i <= n; i++ {
		if i >= 2 {
			fn[i] = fn[i-1] + fn[i-2]
		} else {
			fn[i] = fn[i-1]
		}
	}
	return fn[n]
}
