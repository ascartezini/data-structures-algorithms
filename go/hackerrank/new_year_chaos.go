package hackerrank

import "fmt"

// [2, 1, 5, 3, 4]
// [1, 2, 5, 3, 7, 8, 6, 4]
func minimumBribes(q []int32) {

	var bribeCount int32 = 0
	p1 := int32(1)
	p2 := int32(2)
	p3 := int32(3)
	chaos := false

	for i := 0; i < len(q); i++ {
		// no bribe
		if q[i] == p1 {
			p1 = p2
			p2 = p3
			p3++
			// one bribe
		} else if q[i] == p2 {
			bribeCount++
			p2 = p3
			p3++
			// two bribes
		} else if q[i] == p3 {
			bribeCount += 2
			p3++
			// more than two bribes
		} else {
			chaos = true
			break
		}
	}

	if chaos {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribeCount)
	}

}

// pos := int32(i + 1)
// actualPos := q[i] - pos

// more than one bribe
// if actualPos > 2 {
// 	chaos = true
// 	break
// } else if actualPos > 0 {
// 	c += actualPos
// }
// i++
