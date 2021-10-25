package hackerrank

func jumpingOnClouds(c []int32) int32 {

	minJumps := 0
	i := 0
	j := len(c) - 1

	for i < j {

		// skips if it is a thunderhead cloud
		if c[i] == 1 {
			i++
			continue
		}

		if i+2 < j && c[i+2] == 0 {
			i += 2
		} else {
			i++
		}

		minJumps++
	}

	return int32(minJumps)
}
