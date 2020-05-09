package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {

	var bucket = make([][]int, len(nums)+1)

	var fmap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		fmap[nums[i]]++
	}

	for k, v := range fmap {
		bucket[v] = append(bucket[v], k)
	}

	var res []int

	for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
			k -= len(bucket[i])
		}
	}

	return res

}
