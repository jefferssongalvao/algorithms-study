package main

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, num := range nums[:k] {
		sum += num
	}

	currSum := sum
	for i := k; i < len(nums); i++ {
		currSum += nums[i]
		currSum -= nums[i-k]
		if currSum > sum {
			sum = currSum
		}
	}
	return float64(sum) / float64(k)
}
