package hackerrank

func maxSubsetSum(arr []int32) int32 {
	return maxSubsetSumDynamic(arr)
}

func maxTwo(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func maxThree(a, b, c int32) int32 {
	v := int32(0)

	if a >= b && a >= c {
		v = a
	}
	if b >= a && b >= c {
		v = b
	}
	if c >= a && c >= b {
		v = c
	}

	return v
}

func maxSubsetSumDynamic(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}

	// creates a second slice to store the calculations
	maxSums := make([]int32, len(arr))
	// first value of input slice is assigned to first value of maxSums slice
	// the max value between the first two values of the input slice is stored at position 1 of maxSums
	maxSums[0], maxSums[1] = arr[0], maxTwo(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		maxSums[i] = maxThree(arr[i], maxSums[i-1], arr[i]+maxSums[i-2])
	}

	return maxSums[len(arr)-1]
}
