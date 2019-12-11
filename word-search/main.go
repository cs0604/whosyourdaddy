package word_search

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, x int, y int, substr string) bool {

	if x < 0 || x > len(board)-1 {
		return false
	}
	if y < 0 || y > len(board[0])-1 {
		return false
	}

	if len(substr) == 0 {
		return true
	}

	char := substr[0]

	if board[x][y] == char {
		if len(substr) == 1 {
			return true
		}
		//x,y used
		board[x][y] = 255
		//up,right,down,left
		res := search(board, x-1, y, substr[1:]) || search(board, x, y+1, substr[1:]) || search(board, x+1, y, substr[1:]) || search(board, x, y-1, substr[1:])
		//restore x,y
		board[x][y] = char
		return res
	}

	return false
}
