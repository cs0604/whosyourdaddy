package generate_parenthesess

func generateParenthesis(n int) []string {

	var result []string

	helper("", 0, 0, n, &result)

	return result
}

func helper(res string, left int, right int, target int, result *[]string) {
	if len(res) == target*2 {
		*result = append(*result, res)
		return
	}

	if left < target {
		helper(res+"(", left+1, right, target, result)
	}

	if right < left && right < target {
		helper(res+")", left, right+1, target, result)
	}
}
