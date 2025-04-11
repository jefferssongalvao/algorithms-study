package main

import "fmt"

func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {
		currentChar := chars[read]
		count := 0

		for read < len(chars) && chars[read] == currentChar {
			read++
			count++
		}

		chars[write] = currentChar
		write++

		if count > 1 {
			countStr := fmt.Sprintf("%d", count)
			for i := 0; i < len(countStr); i++ {
				chars[write] = countStr[i]
				write++
			}
		}
	}

	return write
}
