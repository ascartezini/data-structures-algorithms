package main

// O(log(n)) time | O(1) space
func binarySearch(array []int, target int) bool {
	left := 0
	right := len(array) - 1

	for left <= right {
		middle := (left + right) / 2    // finds the middle pointer
		potentialMatch := array[middle] // gets its value

		// if the target value is equal to the middle the value was found and returns true
		if target == potentialMatch {
			return true
			// if the target value is lower than the middle value then right pointer assumes middle pointer -1 thus dividing the array in two
		} else if target < potentialMatch {
			right = middle - 1
			// if the target value is greater than the middle value then left pointer assumes middle pointer +1 thus dividing the array in two
		} else {
			left = middle + 1
		}
	}

	return false
}

// func main() {
// 	array := []int{0, 1, 21, 33, 45, 52, 61, 71, 72, 73}
// 	target := 33

// 	fmt.Println(binarySearch(array, target))
// }
