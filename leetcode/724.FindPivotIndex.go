package main

func pivotIndex(nums []int) int {
	sumLeft, sumRight := 0, 0
	for _, num := range nums {
		sumRight += num
	}

	for i, num := range nums {
		if sumLeft == sumRight-sumLeft-num {
			return i
		}
		sumLeft += num
	}

	return -1
}
