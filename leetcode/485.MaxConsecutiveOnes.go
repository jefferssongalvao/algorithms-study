package main

func findMaxConsecutiveOnes(nums []int) int {
	maxOnes, currentMaxOnes := 0, 0
	for _, num := range nums {
		if num == 1 {
			currentMaxOnes++
		} else {
			currentMaxOnes = 0
		}
		maxOnes = max(currentMaxOnes, maxOnes)
	}
	return maxOnes
}
