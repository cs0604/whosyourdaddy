package largest_number

import (
	"sort"
	"strconv"
)

type IntArr []int

func (intarr IntArr) Less(i, j int) bool {
	a := strconv.Itoa(intarr[i])
	b := strconv.Itoa(intarr[j])
	return a+b < b+a
}

func (intarr IntArr) Swap(i, j int) {
	intarr[i], intarr[j] = intarr[j], intarr[i]
}

func (intarr IntArr) Len() int {
	return len(intarr)
}

func largestNumber(nums []int) string {

	sort.Sort(IntArr(nums))

	if nums[len(nums)-1] == 0 {
		return "0"
	}

	res := ""

	for i := len(nums) - 1; i >= 0; i-- {
		res = res + strconv.Itoa(nums[i])
	}

	return res
}
