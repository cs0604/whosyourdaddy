package majority_element_ii

func majorityElement(nums []int) []int {

	var num1, num2, count1, count2 int

	for i := 0; i < len(nums); i++ {
		if count1 > 0 && num1 == nums[i] {
			count1++
			continue
		}
		if count2 > 0 && num2 == nums[i] {
			count2++
			continue
		}

		if count1 == 0 {

			num1 = nums[i]
			count1++
			continue
		}

		if count2 == 0 {
			num2 = nums[i]
			count2++
			continue
		}

		count1--
		count2--
	}

	count1 = 0
	count2 = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == num1 {
			count1++
		} else if nums[i] == num2 {
			count2++
		}
	}

	var res []int
	if count1 > len(nums)/3 {
		res = append(res, num1)
	}

	if count2 > len(nums)/3 {
		res = append(res, num2)
	}

	return res

}
