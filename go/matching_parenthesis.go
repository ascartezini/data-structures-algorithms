package main

// func main() {
// 	fmt.Println(isMatch("(({{}}))"))
// 	fmt.Println(isMatch("(({{}}))("))
// 	fmt.Println(isMatch("("))
// 	fmt.Println(isMatch("){"))
// }

func isMatch(s string) bool {

	stack := []rune{}

	// if it contains just one character it's invalid
	if len(s) == 1 {
		return false
	}

	for _, v := range s {
		// if it's an opening bracket then add to the stack
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else {
			// returns false is the stack is empty. No openings brackets were added to the stack
			if len(stack) == 0 {
				return false
			}

			potentialMatch := string(stack[len(stack)-1])

			switch v {
			case '}':
				if potentialMatch != "{" {
					return false
				}
			case ']':
				if potentialMatch != "[" {
					return false
				}
			case ')':
				if potentialMatch != "(" {
					return false
				}
			}

			// re-slice the stack by removing the last item. aka pop from stack
			stack = stack[:len(stack)-1]
		}
	}

	// returns true if the stack is empty
	return len(stack) == 0
}
