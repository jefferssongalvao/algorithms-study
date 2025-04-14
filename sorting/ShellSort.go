package main

func shellSort(nums []int) {
	h := 1
	for h < len(nums) {
		h = 3*h + 1
	}

	for h >= 1 {
		h /= 3
		for i := h; i < len(nums); i++ {
			val := nums[i]
			j := i
			for j >= h && nums[j-h] > val {
				nums[j] = nums[j-h]
				j -= h
			}
			nums[j] = val
		}
	}
}
