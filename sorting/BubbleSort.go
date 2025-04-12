package main

import "fmt"

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

func main() {
	nums := []int{8, 4, 1, 12, 4, 3, 6, 7}
	fmt.Println(nums)
	bubbleSort(nums)
	fmt.Println(nums)
}
