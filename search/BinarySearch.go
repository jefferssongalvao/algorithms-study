package main

func binarySearchRecursive(nums []int, num int) int {
	idx := 0
	lastIdx := len(nums) - 1
	for lastIdx >= idx {
		middleIdx := (idx + lastIdx) / 2
		if num > nums[middleIdx] {
			idx = middleIdx + 1
		} else if num < nums[middleIdx] {
			lastIdx = middleIdx - 1
		} else {
			return middleIdx
		}
	}
	return -1
}
