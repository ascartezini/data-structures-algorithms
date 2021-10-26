package hackerrank

import "fmt"

// this challenge is also known as two number sum

func whatFlavors(cost []int32, money int32) {

	visited := map[int32]int32{}

	for i := 0; i < len(cost); i++ {
		// checks if the result of the subtraction money - cost[i] is present in the map
		// if present, prints the index stored for the result and the current index + 1
		if _, present := visited[money-cost[i]]; present {
			fmt.Printf("%d %d\n", visited[money-cost[i]], i+1)
		}

		// stores the visited number and its position in the array
		visited[cost[i]] = int32(i + 1)
	}
}
