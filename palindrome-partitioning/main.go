package palindrome_partitioning

func partition(s string) [][]string {

	var result [][]string
	helper(s, nil, &result)
	return result
}

func helper(str string, tmpRes []string, result *[][]string) {

	if len(str) == 0 {
		res := make([]string, len(tmpRes))
		copy(res, tmpRes)
		*result = append(*result, res)
		return
	}
	for i := 0; i < len(str); i++ {
		if isPalindrom(str[0 : i+1]) {
			tmp := append(tmpRes, str[0:i+1])
			helper(str[i+1:], tmp, result)
		}
	}
}

func isPalindrom(str string) bool {
	if len(str) == 1 {
		return true
	}
	i, j := 0, len(str)-1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
