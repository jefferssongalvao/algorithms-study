package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	minLen := min(len(str1), len(str2))
	for i := minLen; i >= 1; i-- {
		if len(str1)%i != 0 && len(str2)%i != 0 {
			continue
		}

		candidate := str1[:i]
		if len(str2) == minLen {
			candidate = str2[:i]
		}

		if strings.Repeat(candidate, len(str2)/i) == str2 &&
			strings.Repeat(candidate, len(str1)/i) == str1 {
			return candidate
		}
	}
	return ""
}
