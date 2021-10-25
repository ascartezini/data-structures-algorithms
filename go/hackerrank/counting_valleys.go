package hackerrank

func countingValleys(steps int32, path string) int32 {
	stepCounter := 0
	valleys := 0

	for _, v := range path {

		if stepCounter == 0 && string(v) == "D" {
			valleys += 1
		}

		if string(v) == "D" {
			stepCounter--
		} else {
			stepCounter++
		}
	}

	return int32(valleys)
}
