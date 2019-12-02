package palindromic_substrings

func countSubstrings(s string) int {

	var sum int
	for j := 0; j < len(s); j++ {
		for i := 1; i <= len(s)-j; i++ {
			sum += helper(s[j : j+i])
		}
	}

	return sum
}

// abcabc

func helper(cur string) int {
	i := 0
	j := len(cur) - 1

	for i < j && cur[i] == cur[j] {
		i++
		j--
	}

	if i >= j {
		return 1
	}
	return 0
}
