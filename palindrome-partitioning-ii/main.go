package palindrome_partitioning_ii

func minCut(s string) int {

	pd := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		pd[i] = make([]bool, len(s))
	}

	cut := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		cut[i] = i
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || pd[j+1][i-1]) {
				pd[j][i] = true
				// cut[i]=min(cut[i],j==0?0:cut[j]+1)
				cur := 0
				if j > 0 {
					cur = cut[j-1] + 1
				}
				if cut[i] > cur {
					cut[i] = cur
				}
			}
		}
	}

	return cut[len(s)-1]
}
