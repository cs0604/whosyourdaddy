package longest_consecutive_sequence

func longestConsecutive(nums []int) int {

	var vmap = make(map[int]bool, 10)
	for _, v := range nums {
		vmap[v] = true
	}
	var max = 0
	for _, v := range nums {
		cons := 1
		//判断大于v的连续数
		for i := 1; ; i++ {
			if b, ok := vmap[v+i]; ok && b {
				cons++
				vmap[v+i] = false
			} else {
				break
			}
		}

		////判断小于v的连续数
		for i := 1; ; i++ {
			if b, ok := vmap[v-i]; ok && b {
				cons++
				vmap[v-i] = false
			} else {
				break
			}
		}

		if cons > max {
			max = cons
		}
	}

	return max

}
