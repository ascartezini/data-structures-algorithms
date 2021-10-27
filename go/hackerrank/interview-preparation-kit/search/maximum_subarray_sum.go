package hackerrank

// max contiguous sum

func maximumSum(a []int64) int64 {
	maxSum := a[0]
	curSum := maxSum

	for i := 1; i < len(a); i++ {
		curSum = maxTwo(a[i], a[i]+curSum) // is the current value greater than current value + current sum?
		maxSum = maxTwo(curSum, maxSum)
	}

	return maxSum
}

func maxTwo(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
