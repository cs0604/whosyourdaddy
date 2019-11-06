package find_all_duplicates_in_an_array

func findDuplicates(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		k := nums[i]
		if k < 0 {
			k = -k
		}
		if nums[k-1] < 0 {
			res = append(res, k)
		}
		nums[k-1] = -nums[k-1]
	}

	return res
}
