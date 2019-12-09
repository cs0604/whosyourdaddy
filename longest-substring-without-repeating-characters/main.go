package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {

	imap := make(map[int32]int)
	maxlen := 0
	start := 0
	for i, v := range s {

		j, ok := imap[v]
		if ok && start < j+1 {
			start = j + 1
		} else {
			if maxlen < i-start+1 {
				maxlen = i - start + 1
			}
		}
		imap[v] = i
	}

	return maxlen
}
