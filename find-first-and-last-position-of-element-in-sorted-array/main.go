package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {

	low := 0
	high := len(nums) - 1

	//find low
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low > len(nums)-1 || nums[low] != target {
		return []int{-1, -1}
	}

	res := []int{low}

	//find high
	high = len(nums) - 1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	res = append(res, high)

	return res

}
