package partition_equal_subset_sum

func canPartition(nums []int) bool {

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	res := sum / 2

	var dp = make([]bool, res+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := res; j > 0; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[res]

}

// 1 2 3 4 5 6 7

//0 1  2  3
