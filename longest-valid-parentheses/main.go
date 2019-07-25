package longest_valid_parentheses

func longestValidParentheses(s string) int {

	var index, score []int
	score = make([]int, len(s))
	index = make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			index = append(index, i)
		} else {
			if len(index) > 0 {
				//出栈
				top := index[len(index)-1]
				score[top] += 2
				index = index[:len(index)-1]

				//栈顶score + 2
				if len(index) > 0 {
					score[index[len(index)-1]] += score[top]
				}
			}
		}
	}

	//栈里面残留的score = 0
	for _, v := range index {
		score[v] = 0
	}

	//遍历score
	var max, cnt int
	for i := 0; i < len(score); {

		if score[i] > 0 {
			cnt += score[i]
			i += score[i]
			if max < cnt {
				max = cnt
			}
		} else {
			cnt = 0
			i++
		}
	}

	return max
}
