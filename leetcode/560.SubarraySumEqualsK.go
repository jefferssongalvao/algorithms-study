package main

func subarraySum(nums []int, k int) int {
	prefixSum, sum, count := make(map[int]int, len(nums)), 0, 0
	for i, _ := range nums {
		sum += nums[i]
		if sum == k {
			count++
		}
		if freq, ok := prefixSum[sum-k]; ok {
			count += freq
		}
		prefixSum[sum]++
	}
	return count
}
