package main

func subarraySum(nums []int, k int) int {
	prefixSum, sum, count := map[int]int{0: 1}, 0, 0
	for _, num := range nums {
		sum += num
		if freq, ok := prefixSum[sum-k]; ok {
			count += freq
		}
		prefixSum[sum] += 1
	}
	return count
}
