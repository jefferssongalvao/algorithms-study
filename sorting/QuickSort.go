package main

func quickSort(arr []int, startIdx, endIdx int) {
	if startIdx >= endIdx {
		return
	}

	pivot := arr[startIdx+(endIdx-startIdx)/2]

	left := startIdx
	right := endIdx

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quickSort(arr, startIdx, right)
	quickSort(arr, left, endIdx)
}
