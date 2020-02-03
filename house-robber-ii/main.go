package house_robber_ii

func max(args ...int) int {

	ret := -1 << 63

	for _, v := range args {
		if v > ret {
			ret = v
		}
	}
	return ret

}

func calc(nums []int) int {

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

func rob(nums []int) int {

	//处理一下特例
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	a1 := nums[0] + calc(nums[2:len(nums)-1])

	a2 := calc(nums[1:])

	return max(a1, a2)
}
