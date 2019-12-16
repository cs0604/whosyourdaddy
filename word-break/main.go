package word_break

func wordBreak(s string, wordDict []string) bool {

	var dict = make(map[string]bool)

	for _, v := range wordDict {
		dict[v] = true
	}

	return canBreak(s, dict)

}

//leetcode
func canBreak(s string, dict map[string]bool) bool {
	for i := 1; i <= len(s); i++ {
		if dict[s[:i]] == true {

			if i == len(s) {
				return true
			}

			res := canBreak(s[i:], dict)

			if res {
				return res
			}

		}
	}

	return false
}
