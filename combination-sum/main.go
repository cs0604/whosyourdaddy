package combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	var result [][]int

	var queue [][]int

	sort.Ints(candidates)

	//初始化queue, 0:sum  1:index  2... : 组成元素
	for i := 0; i < len(candidates); i++ {
		if candidates[i] == target {
			result = append(result, []int{candidates[i]})
		} else if candidates[i] < target {
			sum := candidates[i]
			tmp := []int{sum, i, candidates[i]}
			queue = append(queue, tmp)
			for {
				if sum+candidates[i] == target {
					//add to result
					tmp = append(tmp, candidates[i])
					result = append(result, tmp[2:])
					break
				} else if sum+candidates[i] > target {
					break
				}
				if sum+candidates[i] < target {
					sum += candidates[i]
					cur := make([]int, len(tmp))
					copy(cur, tmp)
					cur[0] = sum
					cur = append(cur, candidates[i])
					queue = append(queue, cur)
					tmp = cur
				}
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		for i := cur[1] + 1; i < len(candidates); i++ {

			if cur[0]+candidates[i] == target {
				//find one
				cur = append(cur, candidates[i])
				result = append(result, cur[2:])
				break
			} else if cur[0]+candidates[i] < target {
				//in queue
				sum := cur[0] + candidates[i]
				tmp := make([]int, len(cur))
				copy(tmp, cur)
				tmp[0] += candidates[i]
				tmp[1] = i
				tmp = append(tmp, candidates[i])
				queue = append(queue, tmp)

				for {

					if sum+candidates[i] == target {
						//add to result
						tmp = append(tmp, candidates[i])
						result = append(result, tmp[2:])
						break
					} else if sum+candidates[i] > target {
						break
					}
					if sum+candidates[i] < target {
						//copy cur
						sum += candidates[i]
						tmp2 := make([]int, len(tmp))
						copy(tmp2, tmp)
						tmp2[0] = sum
						tmp2 = append(tmp2, candidates[i])
						queue = append(queue, tmp2)
						tmp = tmp2
					}
				}

			} else {
				break
			}
		}
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			break
		}
	}
	return result
}
