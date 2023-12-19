package main

import (
	"fmt"
)

// ToDo cant

func main() {
	fmt.Println(longestPalindrome("babad"), "aba")
	fmt.Println(longestPalindrome("cbbd"), "bb")
	fmt.Println(longestPalindrome("c"), "c")
	fmt.Println(longestPalindrome(""), "")
	fmt.Println(longestPalindrome("xaabcaaa"), "aaa")

}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	left := 0
	right := len(s) - 1

	for left < right {

		if isPalindrome(s, left, right) {
			return s[left:right]
		}
		if isPalindrome(s, left+1, right) {
			return s[left:right]
		}
		left++
	}

	return ""
}

func isPalindrome(s string, left, right int) bool {
	for left < right-1 {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}
