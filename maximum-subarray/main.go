package maximum_subarray

func maxSubArray(nums []int) int {
	sum := 0
	maxsum := -1 << 63
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if maxsum < sum {
			maxsum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxsum
}
