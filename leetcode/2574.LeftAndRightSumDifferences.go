package main

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func leftRightDifference(nums []int) []int {
	rightSum, leftSum := 0, 0
	for _, n := range nums {
		rightSum += n
	}

	for i, n := range nums {
		rightSum -= n
		sumAbs := abs(leftSum - rightSum)
		leftSum += n
		nums[i] = sumAbs
	}
	return nums
}
