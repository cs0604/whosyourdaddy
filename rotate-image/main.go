package rotate_image

/**
1 2 3        7 4 1
4 5 6   =>   8 5 2
7 8 9        9 6 3


9 4 7
8 5 2
3 6 1
*/
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	N := len(matrix)
	for i := 0; i < N/2; i++ {
		for j := i; j < N-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[N-j-1][i]
			matrix[N-j-1][i] = matrix[N-i-1][N-j-1]
			matrix[N-i-1][N-j-1] = matrix[j][N-i-1]
			matrix[j][N-i-1] = tmp
		}
	}
}
