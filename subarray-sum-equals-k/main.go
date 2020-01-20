package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	var count = 0
	var preMap = make(map[int]int, len(nums))
	var sum = 0
	preMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		find := sum - k
		if cnt, ok := preMap[find]; ok {
			count += cnt
		}
		preMap[sum]++
	}
	return count
}
