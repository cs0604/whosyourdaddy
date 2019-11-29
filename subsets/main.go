package subsets

import "sort"

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var dst []int
	for i := 0; i <= len(nums); i++ {
		calc(nums, &dst, 0, i, &result)
	}
	return result
}

func calc(nums []int, dst *[]int, startIndex int, target int, result *[][]int) {
	if len(*dst) == target {
		tmp := make([]int, len(*dst))
		copy(tmp, *dst)
		*result = append(*result, tmp)
		return
	}
	for i := startIndex; i < len(nums); i++ {
		*dst = append(*dst, nums[i])
		calc(nums, dst, i+1, target, result)
		*dst = (*dst)[:len(*dst)-1]
	}
}
