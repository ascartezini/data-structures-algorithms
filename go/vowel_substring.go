package main

import (
	"fmt"
	"time"
)

func findSubstring(s string, k int32) string {

	var word string
	var wordAux string
	var maxCount int
	var currentCount int
	vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}

	for i := 0; i <= len(s)-1; i++ {
		char := string(s[i])
		wordAux += char

		// if the current char is a vowel then increase the counter
		if _, present := vowels[char]; present {
			currentCount++
		}

		// if the length of wordAux is of size ok k
		if int32(len(wordAux))/k == 1 {

			// if current count is greater than the maximum counter then update the maxCount
			if currentCount > maxCount {
				word = wordAux
				maxCount = currentCount
			}

			// if the char that is about to be removed from the string is a vowel then the counter is decreased by 1
			if _, present := vowels[string(wordAux[0])]; present {
				currentCount--
			}

			// discards the first char
			wordAux = wordAux[1:]
		}
	}

	if maxCount == 0 {
		return "Not found!"
	}

	return word
}

func main() {
	start := time.Now()
	fmt.Println(findSubstring("andregregorioscartezini", 5))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
