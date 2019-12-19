package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {

	var result []int

	var parr [26]int

	for i := 0; i < len(p); i++ {
		parr[p[i]-'a']++
	}

	diff := len(p)

	i := 0
	j := 0

	//初始化，先把j移动到合适的位置，令j-i+1=len(p)
	for j < len(s) && j-i+1 <= len(p) {
		parr[s[j]-'a']--
		if parr[s[j]-'a'] >= 0 {
			diff--
		}
		j++
	}
	if diff == 0 {
		result = append(result, i)
	}
	//i和j开始往后移
	for j < len(s) {
		parr[s[i]-'a']++
		if parr[s[i]-'a'] > 0 {
			diff++
		}
		i++
		parr[s[j]-'a']--
		if parr[s[j]-'a'] >= 0 {
			diff--
		}
		j++
		if diff == 0 {
			result = append(result, i)
		}
	}
	return result
}
