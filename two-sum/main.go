package two_sum

func twoSum(nums []int, target int) []int {
	var imap = make(map[int][]int, 10)
	//把索引放进去
	for i := 0; i < len(nums); i++ {
		index, _ := imap[nums[i]]
		index = append(index, i)
		imap[nums[i]] = index
	}
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		index, ok := imap[need]
		if ok {
			for j := 0; j < len(index); j++ {
				if index[j] != i {
					//have found
					return []int{i, index[j]}
				}
			}
		}
	}
	return nil
}
