package unique_binary_search_trees

func numTrees(n int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return calc(nums)
}

func calc(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	if len(nums) == 1 {
		return 1
	}

	cnt := 0
	for i := 0; i < len(nums)/2; i++ {
		if i == len(nums)-1 {
			cnt += calc(nums[0:i])
		} else {
			cnt += calc(nums[0:i]) * calc(nums[i+1:])
		}
	}
	if len(nums)%2 == 1 {
		mid := calc(nums[0 : len(nums)/2])
		cnt += cnt + mid*mid
	} else {
		cnt += cnt
	}
	return cnt
}
