package hackerrank

func deleteChar(mapA map[string]int32, mapB map[string]int32) int32 {

	var charDeleted int32
	for letterA := range mapA {
		// if the current character of mapA is present in mapB then delete it from A and increase the counter
		if _, present := mapB[letterA]; !present {
			charDeleted += mapA[letterA]
			delete(mapA, letterA)
		} else {
			countLetterA := mapA[letterA]
			countLetterB := mapB[letterA]
			var deleted int32 = 0
			// if the same letter is more frequent in mapA than mapB
			if countLetterA > countLetterB {
				// calculates the difference between the two
				deleted = countLetterA - countLetterB
				// updates the frequency in mapA to the frequency in mapB
				mapA[letterA] = countLetterB
			} else if countLetterA < countLetterB {
				deleted = countLetterB - countLetterA
				mapB[letterA] = countLetterA
			}

			charDeleted += deleted
		}
	}

	return charDeleted
}

func makeAnagram(a string, b string) int32 {

	var charDeleted int32 = 0
	mapA := map[string]int32{}
	mapB := map[string]int32{}

	// stores the frequency of each letter in a
	for _, v := range a {
		mapA[string(v)]++
	}

	// stores the frequency of each letter in b
	for _, v := range b {
		mapB[string(v)]++
	}

	charDeleted = deleteChar(mapA, mapB)
	charDeleted += deleteChar(mapB, mapA)

	return charDeleted
}
