package main

import (
	"fmt"
)

// ToDo cant

func main() {
	//fmt.Println(longestPalindrome("babad"), "bab | aba")
	fmt.Println(longestPalindrome("cbbd"), "bb")
	//fmt.Println(longestPalindrome("c"), "c")
	//fmt.Println(longestPalindrome(""), "")

}

func longestPalindrome(s string) string {
	for start := 0; start < len(s); start++ {
		for end := len(s) - 1; end >= 0; end-- {

			if isPalindrome(s, start, end) {
				return s[start : start+end]
			}
		}
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
