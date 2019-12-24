package longest_increasing_subsequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var dp = make([]int, len(nums))

	var curIndex int

	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {

		//如果超过最大值，直接插入到尾端，curIndex++
		if dp[curIndex] < nums[i] {
			curIndex++
			dp[curIndex] = nums[i]
		} else {
			//否则只需要更新dp[i],不影响最终结果
			low := 0
			high := curIndex

			for low <= high {
				mid := (low + high) >> 1
				if nums[i] <= dp[mid] {
					high = mid - 1
				} else if nums[i] > dp[mid] {
					low = mid + 1
				}
			}

			dp[low] = nums[i]
		}
	}

	return curIndex + 1
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	var dp = make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}
