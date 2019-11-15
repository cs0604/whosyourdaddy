package three_sum

import "sort"

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	var res [][]int

	var imap = make(map[int][]int, 10)
	for i := 0; i < len(nums); i++ {
		index, _ := imap[nums[i]]
		index = append(index, i)
		imap[nums[i]] = index
	}

	for i := 0; i < len(nums); i++ {
		//大于0就可以结束了，因为已经排好序
		if nums[i] > 0 {
			break
		}
		//相同的数字不要重复算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {

			//相同的数字不要重复算
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			a := nums[i]
			b := nums[j]
			c := -(a + b)

			index, ok := imap[c]

			if ok {
				for k := 0; k < len(index); k++ {
					if index[k] > j {
						res = append(res, []int{a, b, c})
						break
					}
				}
			}
		}
	}

	return res
}

func threeSum2(nums []int) [][]int {

	sort.Ints(nums)

	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		//大于0就可以结束了，因为已经排好序
		if nums[i] > 0 {
			break
		}
		//相同的数字不要重复算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[low] + nums[high] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[low], nums[high]})
				//找到了一个，需要把low和high往中间靠，相同的排除掉
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if sum < 0 {
				low++
			} else {
				high--
			}
		}

	}

	return res
}

func threeSum3(nums []int) [][]int {
	out := make([][]int, 0)
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			total := nums[i] + nums[l] + nums[r]
			if total < 0 {
				l++
			} else if total > 0 {
				r--
			} else {
				out = append(out, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return out
}
