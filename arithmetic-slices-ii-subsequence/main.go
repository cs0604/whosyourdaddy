package main

func main() {

}

func checkArithmeticSlice(A []int) bool {
	if len(A) < 3 {
		return false
	}
	gap := A[1] - A[0]
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] != gap {
			return false
		}
	}
	return true
}
