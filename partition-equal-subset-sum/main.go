package partition_equal_subset_sum

func canPartition(nums []int) bool {

	sum := 0

	var base [101]int

	for i := 0; i < len(nums); i++ {
		base[nums[i]]++
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	res := sum / 2

	var dp = make([]bool, res+1)
	dp[0] = true

	var tmp = []int{0}

	for i := 0; i < len(tmp); i++ {
		for j := 1; j < 101; j++ {
			for base[j] > 0 {

			}
		}
	}

	return dp[res]

}

// 1 2 3 4 5 6 7

//0 1  2  3
