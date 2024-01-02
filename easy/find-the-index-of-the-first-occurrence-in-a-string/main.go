package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"), "0")
	fmt.Println(strStr("leetcode", "leeto"), "-1")
}

func strStr(haystack string, needle string) int {

	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:len(needle)] == needle {
			return i
		}
	}

	return -1
}
