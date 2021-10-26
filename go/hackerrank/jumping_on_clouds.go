package hackerrank

func jumpingOnClouds(c []int32) int32 {

	minJumps := 0
	i := 0
	j := len(c) - 1

	for i < j {
		// checks if two positions ahead is equal to zero (cumulus cloud)
		if i+2 <= j && c[i+2] == 0 {
			// move the pointer two positions ahead
			i += 2
			// count one more jump
			minJumps++
			// continue to the next iteration
			continue
			// checks if one position is equal to zero (cumulus cloud)
		} else if i+1 <= j && c[i+1] == 0 {
			// move the pointer one position ahead
			i++
			// count one more jump
			minJumps++
			// continue to the next iteration
			continue
		}
	}

	return int32(minJumps)
}
