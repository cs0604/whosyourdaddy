package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {

	var indegree = make([]int, numCourses)

	var graph = make([][]int, numCourses)

	//build graph and indegree
	for i := 0; i < len(prerequisites); i++ {
		x := prerequisites[i][0]
		y := prerequisites[i][1]
		indegree[x]++
		graph[y] = append(graph[y], x)
	}

	//每次从graph中找一个indegree为0的节点开始处理，找不到就说明有环
	//bfs
	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	visit := 0

	for len(queue) > 0 {
		cur := queue[0]
		visit++
		queue = queue[1:]
		for j := 0; j < len(graph[cur]); j++ {
			pos := graph[cur][j]
			indegree[pos]--
			if indegree[pos] == 0 {
				queue = append(queue, pos)
			}
		}
	}

	return visit == numCourses

}
