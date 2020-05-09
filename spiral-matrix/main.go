package spiral_matrix

func spiralOrder(matrix [][]int) []int {

	M := len(matrix)
	if M == 0 {
		return nil
	}
	N := len(matrix[0])

	var res []int
	//M=3,N=4,i=1
	for i := 0; i <= M/2; i++ {

		right := false
		for j := i; j < N-i; j++ {
			res = append(res, matrix[i][j])
			right = true
		}

		if right == false {
			break
		}

		down := false

		for j := i + 1; j < M-i; j++ {
			res = append(res, matrix[j][N-i-1])
			down = true
		}

		if down == false {
			break
		}

		left := false
		for j := N - i - 2; j >= i; j-- {
			res = append(res, matrix[M-i-1][j])
			left = true
		}

		if left == false {
			break
		}

		up := false
		for j := M - i - 2; j >= i+1; j-- {
			res = append(res, matrix[j][i])
			up = true
		}

		if up == false {
			break
		}

	}

	return res

}

func spiralOrder2(matrix [][]int) []int {

	M := len(matrix)
	if M == 0 {
		return nil
	}
	N := len(matrix[0])

	res := []int{}

	top := 0
	bottom := M - 1
	left := 0
	right := N - 1

	for {

		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--

		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++

		if left > right {
			break
		}
	}

	return res

}
