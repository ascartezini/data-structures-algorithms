package hackerrank

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var ap int32
	var bp int32

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			ap++
		}

		if b[i] > a[i] {
			bp++
		}
	}

	return []int32{ap, bp}

}

// func main() {
// 	println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))
// }
