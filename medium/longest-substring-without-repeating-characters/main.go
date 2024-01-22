package main

import (
	"fmt"
)

func main() {
	//fmt.Println(lengthOfLongestSubstring("a"), "1")
	fmt.Println(lengthOfLongestSubstring("aa"), "1")
	fmt.Println(lengthOfLongestSubstring("aaa"), "1")
	fmt.Println(lengthOfLongestSubstring("ab"), "2")
	fmt.Println(lengthOfLongestSubstring("aab"), "2")
	fmt.Println(lengthOfLongestSubstring("dvdf"), "3")
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left := 0
	res := 0
	cache := make(map[string]int)

	for j := 0; j < len(s); j++ {
		v := string(s[j])

		if idx, ok := cache[v]; ok {
			left = min(j, idx)
		}

		res = max(res, j-left+1)
		cache[v] = j + 1

	}

	return res
}
