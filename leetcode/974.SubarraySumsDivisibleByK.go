package main

// func subarraysDivByK(nums []int, k int) int {
// 	count, prefixSum, prefixMap := 0, 0, make(map[int]int, min(k+1, len(nums)))
// 	prefixMap[0] = 1
// 	for idx := 0; idx < len(nums); idx++ {
// 		prefixSum += nums[idx]
// 		mod := ((prefixSum % k) + k) % k
// 		if prefixIdx, present := prefixMap[mod]; present {
// 			count += prefixIdx
// 		}
// 		prefixMap[mod]++
// 	}
// 	return count
// }

func subarraysDivByK(nums []int, k int) int {
	arr := make([]int, k)
	arr[0]++
	s := 0
	for _, v := range nums {
		s = ((s+v)%k + k) % k
		arr[s]++
	}
	res := 0
	for _, v := range arr {
		if v == 0 {
			continue
		}
		res += (v - 1) * v / 2
	}
	return res
}

func main() {
	nums := []int{4, 5, 0, -2, -3, 1}
	k := 5
	result := subarraysDivByK(nums, k)
	println(result) // Output: 7
}
