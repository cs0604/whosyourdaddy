package queue_reconstruction_by_height

import "sort"

type Queue [][]int

func (q Queue) Len() int {
	return len([][]int(q))
}
func (q Queue) Less(i, j int) bool {
	return q[i][0] < q[j][0]
}
func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func reconstructQueue(people [][]int) [][]int {

	var res = make([][]int, len(people))

	//排序
	sort.Sort(Queue(people))

	for i := 0; i < len(people); i++ {
		gap := people[i][1]
		for j := 0; j < len(res); j++ {
			if res[j] == nil {
				if gap == 0 {
					//找到空位就插
					res[j] = people[i]
					break
				}
				gap--
			} else if people[i][0] <= res[j][0] {
				gap--
			}
		}
	}
	return res
}
