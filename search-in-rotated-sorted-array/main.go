package search_in_rotated_sorted_array

func search(nums []int, target int) int {

	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return mid
		}
		if low == mid {
			return bsearch(nums, mid+1, high, target)
		}
		if nums[low] < nums[mid] {
			if target <= nums[mid] && target >= nums[low] {
				return bsearch(nums, low, mid, target)
			}
			low = mid + 1
		} else {
			if target <= nums[high] && target >= nums[mid] {
				return bsearch(nums, mid, high, target)
			}
			high = mid - 1
		}
	}

	return -1

}

//45123
func bsearch(nums []int, low int, high int, target int) int {
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] == target {
			return mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

//1234567
//5671234
