package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {

	var invalid = make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		invalid[i] = make([]bool, numCourses)
	}

	for i := 0; i < len(prerequisites); i++ {
		x := prerequisites[i][0]
		y := prerequisites[i][1]

		if invalid[x][y] {
			return false
		}

		invalid[y][x] = true

		for j := 0; j < numCourses; j++ {
			if invalid[x][j] {
				if j == y {
					return false
				}
				invalid[y][j] = true
			}
		}

	}

	return true

}
