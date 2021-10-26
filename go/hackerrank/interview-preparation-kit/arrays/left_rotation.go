package hackerrank

// [1, 2, 3, 4, 5], 4
// [2, 3, 4, 5, 1]
// [3, 4, 5, 1, 2]
// [4, 5, 1, 2, 3]
// [5, 1, 2, 3, 4]
func rotLeft(a []int32, d int32) []int32 {
	// Write your code here

	left := a[0:d]
	right := a[d:]

	return append(right, left...)
}
