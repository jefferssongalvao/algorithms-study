package main

import "sort"

func sortedSquares(nums []int) []int {
	for i, _ := range nums {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}
