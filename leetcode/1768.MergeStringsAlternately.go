package main

func mergeAlternately(word1 string, word2 string) string {
	newWord := ""
	for i := 0; i < max(len(word1), len(word2)); i++ {
		if i < len(word1) {
			newWord += string(word1[i])
		}
		if i < len(word2) {
			newWord += string(word2[i])
		}
	}
	return newWord
}
