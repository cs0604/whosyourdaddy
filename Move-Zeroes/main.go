package Move_Zeroes

// 1 0 3 0 5 0 0 7 8
// 1 3 5 7 8 0 0 0 0

func moveZeroes(nums []int) {

	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
		} else {
			if cnt > 0 {
				nums[i-cnt] = nums[i]
			}
		}
	}

	for i := 0; i < cnt; i++ {
		nums[len(nums)-1-i] = 0
	}

}
