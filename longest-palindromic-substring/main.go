package longest_palindromic_substring

func longestPalindrome(s string) string {

	maxLen := 0
	var res string
	for i := 0; i < len(s); i++ {

		//奇数
		j, k := i-1, i+1
		cnt := 1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			j--
			k++
			cnt += 2
		}
		if cnt > maxLen {
			maxLen = cnt
			res = s[j+1 : k]
		}

		//偶数
		j, k = i, i+1
		cnt = 0
		for j >= 0 && k < len(s) && s[j] == s[k] {
			j--
			k++
			cnt += 2
		}

		if cnt > maxLen {
			maxLen = cnt
			res = s[j+1 : k]
		}
	}

	return res

}
