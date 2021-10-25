package hackerrank

// i   arr                         swap (indices)
// 0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
// 1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
// 2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
// 3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
// 4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
// 5   [1, 2, 3, 4, 5, 6, 7]
func minimumSwaps(arr []int32) int32 {

	var i int32 = 0
	count := 0

	for i < int32(len(arr)) {
		// is the current element in the correct position?
		// yes
		correctIndex := arr[i] - 1
		if correctIndex == i {
			// move to the next element
			i++
			// no
		} else {
			// swap element to its correct position
			arr[i], arr[correctIndex] = arr[correctIndex], arr[i]
			count++
		}
	}

	return int32(count)
}
