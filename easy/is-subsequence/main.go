package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"), true)
	fmt.Println(isSubsequence("abc", "ahbgd"), false)
	fmt.Println(isSubsequence("axc", "ahbgdc"), false)
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	left := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[left] {
			left++
		}
		if len(s) == left {
			return true
		}
	}

	return left == len(s)
}
