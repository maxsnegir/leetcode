package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"), true)
	fmt.Println(isPalindrome("0P"), false)
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left <= right {
		leftVal := rune(s[left])
		rightVal := rune(s[right])

		if !(unicode.IsLetter(leftVal) || unicode.IsDigit(leftVal)) {
			left++
			continue
		}
		if !(unicode.IsLetter(rightVal) || unicode.IsDigit(rightVal)) {
			right--
			continue
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}
