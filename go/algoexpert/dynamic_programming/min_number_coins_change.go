package algoexpert

func minNumberOfCoinsForChange(amount int, coins []int) int {

	// creates a slice with length of amount + 1.
	// here we create an additional position to acommodate value zero (0) at index 0
	numOfCoins := make([]int, amount+1)

	// sets all values as amount + 1. the idea is to set the biggest possible number
	for i := 0; i < len(numOfCoins); i++ {
		numOfCoins[i] = amount + 1
	}

	// set index 0 to value zero (0)
	numOfCoins[0] = 0

	// calculate the min value for each coin
	for _, coin := range coins {
		for n := range numOfCoins {
			// coin has to be lower than the value stored at numOfCoins
			if coin <= n {
				numOfCoins[n] = min(numOfCoins[n], numOfCoins[n-coin]+1)
			}
		}
	}

	if numOfCoins[amount] != amount+1 {
		return numOfCoins[amount]
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
