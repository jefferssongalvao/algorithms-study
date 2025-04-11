package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	prev := 1
	for i, num := range nums {
		answer[i] = prev
		prev *= num
	}

	post := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= post
		post *= nums[i]
	}

	return answer
}
