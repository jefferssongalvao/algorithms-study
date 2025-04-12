package main

func getAverages(nums []int, k int) []int {
	n, avgs, windowSize := len(nums), make([]int, len(nums)), 2*k+1
	for i := range avgs {
		avgs[i] = -1
	}

	if windowSize > n {
		return avgs
	}

	sum := 0
	for i := 0; i < windowSize; i++ {
		sum += nums[i]
	}

	avgs[k] = sum / windowSize
	for i := windowSize; i < n; i++ {
		sum += nums[i]
		sum -= nums[i-windowSize]
		avgs[i-k] = sum / windowSize
	}
	return avgs
}
