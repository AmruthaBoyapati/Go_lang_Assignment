package main

import "fmt"

func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	longest := ""
	for i := 0; i < len(s); i++ {
		// Odd length palindrome
		pal1 := expandAroundCenter(s, i, i)
		// Even length palindrome
		pal2 := expandAroundCenter(s, i, i+1)

		// Pick the longer
		if len(pal1) > len(longest) {
			longest = pal1
		}
		if len(pal2) > len(longest) {
			longest = pal2
		}
	}
	return longest
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
