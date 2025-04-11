package main

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	reverseString := []rune(s)
	start := 0
	end := len(reverseString) - 1

	for start < end {
		if strings.ContainsRune(vowels, reverseString[start]) && strings.ContainsRune(vowels, reverseString[end]) {
			reverseString[start], reverseString[end] = reverseString[end], reverseString[start]
			start++
			end--
		} else if strings.ContainsRune(vowels, reverseString[start]) {
			end--
		} else {
			start++
		}
	}
	return string(reverseString)
}
