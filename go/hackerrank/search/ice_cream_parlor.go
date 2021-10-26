package hackerrank

import "fmt"

func whatFlavors(cost []int32, money int32) {

	visited := map[int32]int32{}

	for i := 0; i < len(cost); i++ {
		if _, present := visited[money-cost[i]]; present {
			fmt.Printf("%d %d\n", visited[money-cost[i]], i+1)
		}

		visited[cost[i]] = int32(i + 1)
	}
}
