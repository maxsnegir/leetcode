package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3), 3)
	fmt.Println(maxVowels("abciiiidef", 3), 3)
	fmt.Println(maxVowels("leetcode", 3), 2)
	fmt.Println(maxVowels("weallloveyou", 7), 4)
}

func maxVowels(s string, k int) int {
	vowels := map[string]struct{}{
		"a": struct{}{},
		"e": struct{}{},
		"i": struct{}{},
		"o": struct{}{},
		"u": struct{}{},
	}

	var maxV int

	for i := 0; i < k; i++ {
		if _, ok := vowels[string(s[i])]; ok {
			maxV++
		}
	}

	tmp := maxV
	for i := k; i < len(s); i++ {
		if _, ok := vowels[string(s[i-k])]; ok {
			tmp--
		}
		if _, ok := vowels[string(s[i])]; ok {
			tmp++
		}

		if tmp > maxV {
			maxV = tmp
		}
	}

	return maxV
}
