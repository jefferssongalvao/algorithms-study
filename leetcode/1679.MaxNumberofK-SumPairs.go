package main

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	numsOps, l, r := 0, 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == k {
			numsOps++
			l++
			r--
		} else if sum < k {
			l++
		} else {
			r--
		}
	}
	return numsOps

}
