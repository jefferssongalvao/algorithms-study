package main

func subarraysDivByK(nums []int, k int) int {
	count, prefixSum, prefixMap := 0, 0, make(map[int]int, min(k+1, len(nums)))
	prefixMap[0] = 1
	for idx := 0; idx < len(nums); idx++ {
		prefixSum += nums[idx]
		mod := ((prefixSum % k) + k) % k
		if prefixIdx, present := prefixMap[mod]; present {
			count += prefixIdx
		}
		prefixMap[mod]++
	}
	return count
}
