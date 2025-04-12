package main

func binarySearch(nums []int, num, startIdx, finalIdx int) int {
	if startIdx > finalIdx {
		return -1
	}
	middleIdx := (startIdx + finalIdx) / 2
	if nums[middleIdx] == num {
		return middleIdx
	} else if num > nums[middleIdx] {
		return binarySearch(nums, num, middleIdx+1, finalIdx)
	} else {
		return binarySearch(nums, num, startIdx, middleIdx-1)
	}
}
