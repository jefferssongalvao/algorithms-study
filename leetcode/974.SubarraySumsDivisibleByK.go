package main

func subarraysDivByK(nums []int, k int) int {
	arrayQtdMod, s, res := make([]int, k), 0, 0
	arrayQtdMod[0]++
	for _, v := range nums {
		s = ((s+v)%k + k) % k // evita o mod negativo
		arrayQtdMod[s]++
	}
	for _, qtdMod := range arrayQtdMod {
		if qtdMod == 0 {
			continue
		}
		res += (qtdMod * (qtdMod - 1)) / 2
	}
	return res
}
