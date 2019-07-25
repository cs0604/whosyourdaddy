package wildcard_matching

//记忆化搜索
func isMatch(s string, p string) bool {

	var cache = make([][]int, len(s)+1)
	for k := range cache {
		sub := make([]int, len(p)+1)
		cache[k] = sub
	}

	var match func(i int, j int) bool

	match = func(i int, j int) bool {
		if cache[i][j] == 1 {
			return true
		}
		if cache[i][j] == 2 {
			return false
		}

		if i == len(s) && j == len(p) {
			cache[i][j] = 1
			return true
		}

		if i == len(s) {
			if p[j] != '*' {
				cache[i][j] = 2
				return false
			}
			return match(i, j+1)
		}

		if j == len(p) {
			cache[i][j] = 2
			return false
		}

		if p[j] == '?' {
			res := match(i+1, j+1)
			if res {
				cache[i][j] = 1
			} else {
				cache[i][j] = 2
			}
			return res
		}

		if p[j] == '*' {
			res := match(i+1, j) || match(i, j+1) || match(i+1, j+1)
			if res {
				cache[i][j] = 1
			} else {
				cache[i][j] = 2
			}
			return res
		}

		if s[i] != p[j] {
			cache[i][j] = 2
			return false
		}

		res := match(i+1, j+1)
		if res {
			cache[i][j] = 1
		} else {
			cache[i][j] = 2
		}
		return res

	}

	return match(0, 0)

}

//动态规划
func isMatch2(s string, p string) bool {

	var cache = make([][]bool, len(s)+1)
	for k := range cache {
		sub := make([]bool, len(p)+1)
		cache[k] = sub
	}

	//init
	cache[0][0] = true
	for i := 0; i < len(p); i++ {
		if p[i] == '?' {
			cache[0][i+1] = false
		} else if p[i] == '*' {
			cache[0][i+1] = cache[0][i]
		} else {
			cache[0][i+1] = false
		}
	}
	for i := 0; i < len(s); i++ {
		cache[i+1][0] = false
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] {
				cache[i+1][j+1] = cache[i][j]
			} else if p[j] == '?' {
				cache[i+1][j+1] = cache[i][j]
			} else if p[j] == '*' {
				cache[i+1][j+1] = cache[i][j] || cache[i+1][j] || cache[i][j+1]
			}
		}
	}
	return cache[len(s)][len(p)]
}
