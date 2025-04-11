package main

func longestSubarray(nums []int) int {
	left, maxLength, zeroCount := 0, 0, 0
	for right, num := range nums {
		if num == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxLength = max(maxLength, right-left)
	}

	return maxLength
}
