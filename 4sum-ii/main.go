package _sum_ii

func fourSumCount(A []int, B []int, C []int, D []int) int {

	var tmap = make(map[int]int)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			key := A[i] + B[j]
			tmap[key]++
		}
	}

	res := 0

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sum := C[i] + D[j]
			res += tmap[-sum]
		}
	}

	return res

}
