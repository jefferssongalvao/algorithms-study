package main

func waysToSplitArray(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	sum, count := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		sum += nums[i]
		if sum >= (total - sum) {
			count++
		}
	}
	return count
}
