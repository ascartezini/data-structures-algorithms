package main

import "fmt"

func validAnagram(word1, word2 string) bool {
	// if the length is different then it's not a valid anagram
	if len(word1) != len(word2) {
		return false
	}

	wordMap1 := map[string]int{}
	wordMap2 := map[string]int{}

	// creates an object containing the count for each character for word1
	for _, char := range word1 {
		letter := string(char)
		counter := 0

		if counter, found := wordMap1[letter]; found {
			counter++
		}

		wordMap1[letter] = counter
	}

	// creates an object containing the count for each character for word2
	for _, char := range word2 {
		letter := string(char)
		counter := 0

		if counter, found := wordMap2[letter]; found {
			counter++
		}

		wordMap2[letter] = counter
	}

	// loop through the keys of wordMap1
	for k := range wordMap1 {
		// checks if there is a letter key of wordMap1 in wordMap2
		if _, found := wordMap2[k]; !found {
			return false
		}

		// checks if the count of the letter key of wordMap1 in wordMap2 is the same
		if wordMap1[k] != wordMap2[k] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(validAnagram("god", "dog"))
	fmt.Println(validAnagram("race", "care"))
	fmt.Println(validAnagram("earth", "heart"))
}
