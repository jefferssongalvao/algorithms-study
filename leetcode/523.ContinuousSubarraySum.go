package main

func checkSubarraySum(nums []int, k int) bool {
	prefixSum, prefixMap := 0, make(map[int]int, min(k+1, len(nums)))
	prefixMap[0] = -1
	for idx := range nums {
		prefixSum += nums[idx]
		if prefixIdx, present := prefixMap[prefixSum%k]; present {
			if idx-prefixIdx > 1 {
				return true
			}
		} else {
			prefixMap[prefixSum%k] = idx
		}
	}
	return false
}
