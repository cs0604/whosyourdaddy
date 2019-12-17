package word_break

type trie struct {
	isWord bool
	node   [26]*trie
}

func wordBreak(s string, wordDict []string) bool {

	root := &trie{}

	subres := make(map[string]int)

	for _, word := range wordDict {
		p := root
		for _, v := range word {
			if p.node[v-'a'] == nil {
				p.node[v-'a'] = &trie{}
			}
			p = p.node[v-'a']
		}
		p.isWord = true
	}

	return canBreak(s, root, subres)

}

//leetcode
func canBreak(s string, root *trie, subres map[string]int) bool {

	if subres[s] != 0 {
		return subres[s] > 0
	}

	p := root
	for i := 0; i < len(s); i++ {
		if p.node[s[i]-'a'] != nil {
			p = p.node[s[i]-'a']

			if p.isWord {

				if i == len(s)-1 {
					subres[s] = 1
					return true
				}

				res := canBreak(s[i+1:], root, subres)

				if res {
					subres[s] = 1
					return true
				}
				subres[s] = -1
			}

		} else {
			subres[s] = -1
			return false
		}
	}

	subres[s] = -1
	return false
}
