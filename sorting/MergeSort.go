package main

func mergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	midIdx := len(nums) / 2

	mergeSort(nums[:midIdx])
	mergeSort(nums[midIdx:])

	merge(nums, midIdx)
}

func merge(nums []int, midIdx int) {
	left := make([]int, midIdx)
	copy(left, nums[:midIdx])

	i, j, k := 0, midIdx, 0
	for i < len(left) && j < len(nums) {
		if left[i] <= nums[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = nums[j]
			j++
		}
		k++
	}

	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
}
