package main

func findMiddleIndex(nums []int) int {
	leftSum, rightSum := 0, 0
	for _, num := range nums {
		rightSum += num
	}

	for i := 0; i < len(nums); i++ {
		if rightSum-nums[i] == leftSum {
			return i
		}
		rightSum -= nums[i]
		leftSum += nums[i]
	}
	return -1
}
