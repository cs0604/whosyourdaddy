package increasing_triplet_subsequence

func increasingTriplet(nums []int) bool {

	if len(nums) < 3 {
		return false
	}

	var dp [2]int
	dp[0] = nums[0]
	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[j] {
			j++
			if j == 2 {
				return true
			}
			dp[j] = nums[i]
		} else {
			if nums[i] > dp[0] {
				dp[1] = nums[i]
			} else {
				dp[0] = nums[i]
			}
		}
	}

	return false

}
