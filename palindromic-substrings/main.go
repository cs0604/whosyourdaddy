package palindromic_substrings

func countSubstrings(s string) int {

	var sum int
	for i := 0; i < len(s); i++ {
		sum += helper(s, i)
	}

	return sum
}

// abcabc

func helper(str string, pos int) int {

	cnt := 1

	//回文长度是奇数
	for i, j := pos-1, pos+1; i >= 0 && j < len(str) && str[i] == str[j]; {
		i--
		j++
		cnt++
	}

	//回文长度是偶数
	for i, j := pos, pos+1; i >= 0 && j < len(str) && str[i] == str[j]; {
		i--
		j++
		cnt++
	}

	return cnt
}
