package find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {

	for i := 0; i < len(nums); {

		k := nums[i]

		if k != i+1 && nums[k-1] != k {
			//need swap
			tmp := nums[k-1]
			nums[k-1] = k
			nums[i] = tmp
		} else {
			i++
		}
	}
	var result []int
	for i := 0; i < len(nums); i++ {

		if nums[i] != i+1 {
			result = append(result, i+1)
		}

	}

	return result
}
