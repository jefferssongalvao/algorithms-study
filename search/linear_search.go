package main

func linearSearch(nums []int, num int) int {
	for i, n := range nums {
		if n == num {
			return i
		}
	}
	return -1
}
