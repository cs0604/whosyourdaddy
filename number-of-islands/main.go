package number_of_islands

func numIslands(grid [][]byte) int {

	row := len(grid)

	if row == 0 {
		return 0
	}

	col := len(grid[0])

	if col == 0 {
		return 0
	}

	count := 0

	var visit = make([][]bool, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if !visit[i][j] {
				if grid[i][j] == '1' {
					count++
				}
				helper(grid, i, j, visit)
			}
		}
	}

	return count

}

func helper(grid [][]byte, i int, j int, visit [][]bool) {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 {
		return
	}
	if visit[i][j] {
		return
	}
	visit[i][j] = true
	if grid[i][j] == '1' {
		helper(grid, i, j+1, visit)
		helper(grid, i+1, j, visit)
		helper(grid, i-1, j, visit)
		helper(grid, i, j-1, visit)
	}
}
