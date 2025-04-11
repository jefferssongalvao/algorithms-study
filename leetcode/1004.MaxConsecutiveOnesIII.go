package main

func longestOnes(nums []int, k int) int {
	low, high := 0, 0

	for high < len(nums) {
		if nums[high] == 0 {
			k--
		}
		high++
		if k < 0 {
			if nums[low] == 0 {
				k++
			}
			low++
		}
	}
	return high - low
}
