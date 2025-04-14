package main

func radixSortBits(nums []int, start, end int, mask uint) {
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

	radixSortBits(nums, start, left-1, mask>>1)
	radixSortBits(nums, left, end, mask>>1)
}

// radixSort é a função inicial
func radixSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	// Descobre o bit mais significativo
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	var mask uint = 1
	for mask <= uint(max) {
		mask <<= 1
	}
	mask >>= 1

	radixSortBits(nums, 0, len(nums)-1, mask)
}
