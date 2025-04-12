package main

func twoSum(nums []int, target int) []int {
	prefix := make(map[int]int, 0)
	for i, num := range nums {
		if j, ok := prefix[target-num]; ok {
			return []int{i, j}
		}
		prefix[num] = i
	}
	return []int{}
}
