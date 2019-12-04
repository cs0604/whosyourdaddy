package shortest_unsorted_continuous_subarray

import "sort"

func findUnsortedSubarray(nums []int) int {

	tmp := make([]int, len(nums))
	copy(tmp, nums)

	sort.Ints(nums)

	var left, right int
	for left = 0; left < len(nums); left++ {
		if nums[left] != tmp[left] {
			break
		}
	}
	for right = len(nums) - 1; right > left; right-- {
		if nums[right] != tmp[right] {
			break
		}
	}

	if right-left >= 0 {
		return right - left + 1
	}

	return 0

}

func min(a ...int) int {
	m := 1<<63 - 1
	for _, n := range a {
		if m > n {
			m = n
		}
	}
	return m
}

func max(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if m < num {
			m = num
		}
	}
	return m
}

func findUnsortedSubarray2(nums []int) int {
	var stack []int
	l, r := len(nums), 0
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			l = min(l, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			r = max(r, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	if r-l > 0 {
		return r - l + 1
	} else {
		return 0
	}
}
