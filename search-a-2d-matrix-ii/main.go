package search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var bx, by, ex, ey int //left point， right point

	bx = 0
	by = 0
	ex = len(matrix) - 1
	ey = len(matrix[0]) - 1

	for bx <= ex && by <= ey {
		//trim right column
		found, index := bsearch(matrix, target, bx, by, ex, ey, true)

		if found {
			return true
		}

		if index < by {
			return false
		}

		ey = index
		//trim top row
		found, index = bsearch(matrix, target, bx, by, ex, ey, false)
		if found {
			return true
		}
		if index > ex {
			return false
		}
		bx = index + 1
	}

	return false

}

//如果不存在就返回第一个小于target的index
func bsearch(matrix [][]int, target int, bx, by, ex, ey int, findInRow bool) (bool, int) {
	if findInRow {
		for by <= ey {
			mid := (by + ey) >> 1
			if matrix[bx][mid] == target {
				return true, mid
			}
			if matrix[bx][mid] < target {
				by = mid + 1
			} else {
				ey = mid - 1
			}
		}
		return false, ey
	} else {
		for bx <= ex {
			mid := (bx + ex) >> 1
			if matrix[mid][ey] == target {
				return true, mid
			}
			if matrix[mid][ey] < target {
				bx = mid + 1
			} else {
				ex = mid - 1
			}
		}
		return false, ex
	}

}
