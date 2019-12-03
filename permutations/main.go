package permutations

//1 2 3 -> 2 1 3  -> 2 3 1
//      -> 3 2 1 -> 3 1 2
//      -> 1 3 2
func permute(nums []int) [][]int {
	var result [][]int
	helper(nums, 0, &result)
	return result
}

func helper(cur []int, start int, result *[][]int) {
	if start == len(cur)-1 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	for i := start; i < len(cur); i++ {
		//swap start and i
		cur[start], cur[i] = cur[i], cur[start]
		helper(cur, start+1, result)
		cur[start], cur[i] = cur[i], cur[start]
	}
}
