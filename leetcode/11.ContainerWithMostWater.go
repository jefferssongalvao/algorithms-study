package main

func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1
	for left < right {
		maxArea = max(maxArea, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
			continue
		}
		right--
	}
	return maxArea
}
