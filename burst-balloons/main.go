package main

var dp [502][502]int

func maxCoins(nums []int) int {

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 0
		}
	}
	var cnums []int
	//初始化(首尾都是1）
	cnums = append(cnums, 1)
	cnums = append(cnums, nums...)
	cnums = append(cnums, 1)

	return calc2(cnums, 1, len(cnums)-2)

}

func calc(nums []int, low int, high int) int {

	if low == high {
		return nums[low] * nums[low-1] * nums[low+1]
	}
	if low > high {
		return 0
	}

	if dp[low][high] > 0 {
		return dp[low][high]
	}

	max := 0
	for i := low; i <= high; i++ {
		val := calc(nums, low, i-1) + nums[i]*nums[low-1]*nums[high+1] + calc(nums, i+1, high)
		if max < val {
			max = val
		}
	}

	dp[low][high] = max

	return max
}

func calc2(nums []int, low int, high int) int {

	n := len(nums)
	for k := 0; k < n-2; k++ {
		for left := 1; left < n-1-k; left++ {
			right := left + k
			for i := left; i <= right; i++ {
				tmp := nums[left-1]*nums[i]*nums[right+1] + dp[left][i-1] + dp[i+1][right]
				if tmp > dp[left][right] {
					dp[left][right] = tmp
				}
			}
		}
	}

	return dp[1][n-2]
}
