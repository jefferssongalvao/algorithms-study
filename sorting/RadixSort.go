package main

func radixSort(nums []int, start, end int, mask uint) {
	if start >= end || mask == 0 {
		return
	}

	left, right := start, end
	for {
		for left <= right && (nums[left]&int(mask)) == 0 {
			left++
		}
		for left <= right && (nums[right]&int(mask)) != 0 {
			right--
		}

		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	radixSort(nums, start, left-1, mask>>1)
	radixSort(nums, left, end, mask>>1)
}
