package surrounded_regions

func solve(board [][]byte) {

	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])

	visit := make([][]bool, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if !visit[i][j] {

				visit[i][j] = true

				if board[i][j] == 'O' {
					//bfs
					queue := [][]int{{i, j}}

					backup := [][]int{}

					needChange := true

					for len(queue) > 0 {
						L := len(queue)
						for k := 0; k < L; k++ {

							p := queue[k]
							backup = append(backup, []int{p[0], p[1]})

							if p[0] == 0 || p[0] == row-1 || p[1] == 0 || p[1] == col-1 {
								needChange = false
							}

							if p[0]+1 < row && board[p[0]+1][p[1]] == 'O' && !visit[p[0]+1][p[1]] {
								visit[p[0]+1][p[1]] = true
								queue = append(queue, []int{p[0] + 1, p[1]})
							}

							if p[0]-1 >= 0 && board[p[0]-1][p[1]] == 'O' && !visit[p[0]-1][p[1]] {
								visit[p[0]-1][p[1]] = true
								queue = append(queue, []int{p[0] - 1, p[1]})
							}

							if p[1]-1 >= 0 && board[p[0]][p[1]-1] == 'O' && !visit[p[0]][p[1]-1] {
								visit[p[0]][p[1]-1] = true
								queue = append(queue, []int{p[0], p[1] - 1})
							}

							if p[1]+1 < col && board[p[0]][p[1]+1] == 'O' && !visit[p[0]][p[1]+1] {
								visit[p[0]][p[1]+1] = true
								queue = append(queue, []int{p[0], p[1] + 1})
							}
						}

						queue = queue[L:]
					}

					if needChange {
						for _, p := range backup {
							board[p[0]][p[1]] = 'X'
						}
					}
				}
			}

		}
	}

}
