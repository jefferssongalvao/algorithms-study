package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	debug.SetMemoryLimit(8)
}

func subarraySum(nums []int, k int) int {
	prefixSum, sum, count := make(map[int]int), 0, 0
	prefixSum[0] = 1
	for _, v := range nums {
		sum += v
		remSum := sum - k
		if value, ok := prefixSum[remSum]; ok {
			count += value
		}
		prefixSum[sum] += 1
	}
	return count
}

func main() {
	nums := []int{1, 2, 1, 2, 1}
	k := 4
	expected := 4
	result := subarraySum(nums, k)
	fmt.Println(expected == result, expected, result)
}
