package main

func lengthOfLongestSubstring(s string) int {
	prefix := make(map[rune]int, 0)
	maxLen, start := 0, 0
	for idx, c := range s {
		if lastIdx, ok := prefix[c]; ok && lastIdx >= start {
			start = lastIdx + 1
		}
		prefix[c] = idx
		proxIdx := idx - start + 1
		if proxIdx > maxLen {
			maxLen = proxIdx
		}
	}
	return maxLen
}
