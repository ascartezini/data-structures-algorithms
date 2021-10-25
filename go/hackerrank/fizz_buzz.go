package hackerrank

import "fmt"

func fizzBuzz(n int32) {
	// Write your code here

	for i := 1; int32(i) <= n; i++ {
		var r string
		if i%3 == 0 {
			r = "Fizz"
		}

		if i%5 == 0 {
			r += "Buzz"
		}

		if len(r) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(r)
		}
	}

}

// func main() {
// 	fizzBuzz(15)
// }
