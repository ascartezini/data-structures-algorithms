package hackerrank

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	pairs := map[int32]int32{}

	for _, v := range ar {
		pairs[v] += 1
	}

	pairCount := 0

	for _, v := range pairs {
		pairCount += int(v / 2)
	}

	return int32(pairCount)
}
