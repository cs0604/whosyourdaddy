package first_missing_positive

func firstMissingPositive(nums []int) int {

	for i, v := range nums {
		//如果当前数不在当前位置，需要移动到它所属的位置上去
		for v > 0 && v-1 < len(nums) && nums[v-1] != v {
			tmp := nums[v-1]
			nums[v-1] = v
			nums[i] = tmp
			v = tmp
		}
	}
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return i + 1
}
