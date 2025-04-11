package main

import "strings"

func maxVowels(s string, k int) int {
	vowels := "aeiou"
	vowelCount, maxVowelCount := 0, 0

	for i := range s[:k] {
		if strings.Contains(vowels, string(s[i])) {
			vowelCount++
		}
	}

	maxVowelCount = vowelCount
	for i := k; i < len(s); i++ {
		if strings.Contains(vowels, string(s[i])) {
			vowelCount++
		}
		if strings.Contains(vowels, string(s[i-k])) {
			vowelCount--
		}
		maxVowelCount = max(vowelCount, maxVowelCount)
	}

	return maxVowelCount
}
