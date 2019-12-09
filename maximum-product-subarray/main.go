package maximum_product_subarray

func maxProduct1(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	max := -1 << 31

	var dp = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			dp[j] = dp[j] * nums[i]
			if dp[j] > max {
				max = dp[j]
			}
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	//2 -3 -4

	var dp = [2]int{0, 0}

	if nums[0] > 0 {
		dp[0] = nums[0]
	} else {
		dp[1] = nums[0]
	}
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			dp[0] = 0
			dp[1] = 0
		} else if nums[i] > 0 {
			dp[0], dp[1] = max(nums[i], dp[0]*nums[i]), nums[i]*dp[1]
		} else {
			dp[0], dp[1] = nums[i]*dp[1], min(nums[i]*dp[0], nums[i])
		}
		if res < dp[0] {
			res = dp[0]
		}
	}

	return res
}
