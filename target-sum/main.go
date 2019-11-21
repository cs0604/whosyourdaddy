package target_sum

//F(S,i) = F(S-nums[i],i+1) + F(S+nums[i], i+1)

func findTargetSumWays(nums []int, S int) int {
	var fmap = make(map[int]int)
	fmap[0] = 1
	for i := 0; i < len(nums); i++ {
		nmap := make(map[int]int)
		for k, v := range fmap {
			nmap[k+nums[i]] = nmap[k+nums[i]] + v
			nmap[k-nums[i]] = nmap[k-nums[i]] + v
		}
		fmap = nmap
	}
	return fmap[S]
}
