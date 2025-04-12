package main

func insertionSort(nums []int) {
	for i := range nums {
		nInsert := nums[i]
		j := i - 1
		for j >= 0 && nInsert < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = nInsert
	}
}
