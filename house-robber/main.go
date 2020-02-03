package house_robber

func max(args ...int) int {

	ret := -1 << 63

	for _, v := range args {
		if v > ret {
			ret = v
		}
	}
	return ret

}

func rob(nums []int) int {

	var dp = make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {

		if i == 0 {
			dp[i+1] = nums[i]
		} else {
			dp[i+1] = max(dp[i], dp[i-1]+nums[i])
		}

	}

	return dp[len(nums)]

}
