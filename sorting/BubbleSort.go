package main

func bubbleSort(nums []int) {
	for i := range nums {
		swapped := false
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

}
