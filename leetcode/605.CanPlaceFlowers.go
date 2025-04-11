package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	total := 0
	for j := 0; j < len(flowerbed); j++ {
		if flowerbed[j] == 0 {
			left := j == 0 || flowerbed[j-1] == 0
			right := j == len(flowerbed)-1 || flowerbed[j+1] == 0
			if left && right {
				flowerbed[j] = 1
				total++
			}
		}
	}
	return total >= n
}
