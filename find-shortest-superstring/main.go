package find_shortest_superstring

import "strings"

func shortestSuperstring(A []string) string {

	var cacheRes [][]string
	var cacheCnt [][]int

	max := len(A)

	cacheRes = make([][]string, 1<<uint(max))
	for k := range cacheRes {
		cacheRes[k] = make([]string, max)
	}
	cacheCnt = make([][]int, 1<<uint(max))
	for k := range cacheCnt {
		cacheCnt[k] = make([]int, max)
	}

	//初始化 假设三个元素
	//fn(0x000,0) = A[0],0
	//fn(0x000,1) = A[1],0
	//fn(0x000,2) = A[2],0

	//求 max { fn(0x011,0), fn(0x101,1), fn(0x110,2) }
	// fn(0x011,0) ====>
	//      fn(0x001,1)+dis(1,0)
	//      fn(0x010,2)+dis(2,0)
	// fn(0x001,1) =======>
	//      fn(0x000,2)+dis(2,1)
	for i := 0; i < max; i++ {
		cacheCnt[0][i] = 0
		cacheRes[0][i] = A[i]
	}

	var fn func(s, k int) (string, int)

	fn = func(s, k int) (string, int) {
		if cacheCnt[s][k] > 0 || cacheRes[s][k] != "" {
			return cacheRes[s][k], cacheCnt[s][k]
		}
		var res string
		var cnt = -1

		for i := 0; i < max; i++ {
			if s&(1<<uint(max-i-1)) != 0 {
				ts := s - (1 << uint(max-i-1))
				tres, tcnt := fn(ts, i)
				tlen := dis(A[i], A[k])
				tcnt += tlen
				tres += A[k][tlen:]

				if tcnt > cnt {
					cnt = tcnt
					res = tres
				}
			}
		}
		cacheRes[s][k] = res
		cacheCnt[s][k] = cnt
		return res, cnt
	}

	var res string
	var cnt = -1
	s := (1 << uint(max)) - 1
	for i := 0; i < max; i++ {
		ts := s - (1 << uint(max-i-1))
		tres, tcnt := fn(ts, i)
		if tcnt > cnt {
			cnt = tcnt
			res = tres
		}
	}

	return res
}

//abcdg  cdefg
func dis(a, b string) int {
	for i := 0; i < len(a); i++ {
		suffix := a[i:]
		if strings.HasPrefix(b, suffix) {
			return len(suffix)
		}
	}
	return 0
}
