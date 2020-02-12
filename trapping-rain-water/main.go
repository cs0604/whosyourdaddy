package trapping_rain_water

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

func trap(height []int) int {

	var leftMax = make([]int, len(height))
	var rightMax = make([]int, len(height))

	leftMax[0] = 0
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}
	rightMax[len(height)-1] = 0
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		if leftMax[i] > height[i] && rightMax[i] > height[i] {
			sum += min(leftMax[i], rightMax[i]) - height[i]
		}
	}

	return sum

}
