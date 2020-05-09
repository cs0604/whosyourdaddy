package word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {

	used := make([]bool, len(wordList))

	nextMatch := make([][]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if checkWord(wordList[i], wordList[j]) {
				nextMatch[i] = append(nextMatch[i], j)
				nextMatch[j] = append(nextMatch[j], i)
			}
		}
	}

	var queue []int

	for i := 0; i < len(wordList); i++ {
		if checkWord(beginWord, wordList[i]) {
			queue = append(queue, i)
			used[i] = true
		}
	}

	cnt := 1

	for len(queue) > 0 {
		L := len(queue)
		cnt++
		for i := 0; i < L; i++ {
			wordIndex := queue[i]
			if wordList[wordIndex] == endWord {
				return cnt
			}
			for j := 0; j < len(nextMatch[wordIndex]); j++ {
				n := nextMatch[wordIndex][j]
				if !used[n] {
					used[n] = true
					queue = append(queue, n)
				}
			}
		}
		queue = queue[L:]
	}

	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {

	used := make([]bool, len(wordList))

	nextMatch := make(map[string][]int)

	for i := 0; i < len(wordList); i++ {

		for j := 0; j < len(wordList[i]); j++ {

			tmp := wordList[i][0:j] + "*" + wordList[i][j+1:]

			arr := nextMatch[tmp]

			nextMatch[tmp] = append(arr, i)

		}

	}

	var queue []int

	for i := 0; i < len(wordList); i++ {
		if checkWord(beginWord, wordList[i]) {
			queue = append(queue, i)
			used[i] = true
		}
	}

	cnt := 1

	wordLen := len(endWord)

	for len(queue) > 0 {
		L := len(queue)
		cnt++
		for i := 0; i < L; i++ {
			wordIndex := queue[i]
			if wordList[wordIndex] == endWord {
				return cnt
			}
			for k := 0; k < wordLen; k++ {
				tmp := wordList[wordIndex][0:k] + "*" + wordList[wordIndex][k+1:]
				for j := 0; j < len(nextMatch[tmp]); j++ {
					n := nextMatch[tmp][j]
					if !used[n] {
						used[n] = true
						queue = append(queue, n)
					}
				}
			}
		}
		queue = queue[L:]
	}

	return 0
}

func checkWord(word1, word2 string) bool {
	cnt := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}
	return true
}
