package main

func containsDuplicate(nums []int) bool {
	prefix := make(map[int]struct{}, len(nums))
	for idx := range nums {
		if _, ok := prefix[nums[idx]]; ok {
			return true
		}
		prefix[nums[idx]] = struct{}{}
	}
	return false
}
