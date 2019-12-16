package merge_intervals

import "sort"

type IntSlice [][]int

func (s IntSlice) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Len() int {
	return len([][]int(s))
}

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return nil
	}

	sort.Sort(IntSlice(intervals))

	var res [][]int

	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]

		if pre[0] <= cur[0] && pre[1] >= cur[0] {
			//merge
			tmp := []int{pre[0], pre[1]}
			if pre[1] < cur[1] {
				tmp[1] = cur[1]
			}
			pre = tmp
		} else {
			res = append(res, pre)
			pre = cur
		}
	}

	res = append(res, pre)

	return res
}
