package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ch := range s {
		// If it's a closing bracket
		if open, ok := brackets[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			stack = append(stack, ch) // push opening bracket
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))      // true
	fmt.Println(isValid("()[]{}"))  // true
	fmt.Println(isValid("(]"))      // false
	fmt.Println(isValid("([)]"))    // false
	fmt.Println(isValid("{[]}"))    // true
}
