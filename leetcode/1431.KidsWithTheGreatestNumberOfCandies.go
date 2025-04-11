package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for i := 0; i < len(candies); i++ {
		maxCandies = max(candies[i], maxCandies)
	}

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= maxCandies
	}
	return result
}
