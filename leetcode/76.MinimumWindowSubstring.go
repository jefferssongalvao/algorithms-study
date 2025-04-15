package main

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	need := [128]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	start, minLen := 0, len(s)+1
	count := len(t)

	for right < len(s) {
		if need[s[right]] > 0 {
			count--
		}
		need[s[right]]--
		right++

		for count == 0 {
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			need[s[left]]++
			if need[s[left]] > 0 {
				count++
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}
