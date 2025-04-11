package main

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			return true
		}
	}
	return false
}
