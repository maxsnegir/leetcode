package main

import (
	"fmt"
)

func main() {

	fmt.Println(mergeAlternately("ab", "pqrs"), "apbqrs")
	fmt.Println(mergeAlternately("abcd", "pq"), "apbqcd")
}

func mergeAlternately(word1 string, word2 string) string {
	res := make([]rune, 0, len(word1)+len(word2))

	left := 0
	right := 0

	for i := 0; i < len(word1)+len(word2); i++ {
		if left < len(word1) {
			res = append(res, rune(word1[left]))
			left++
		}
		if right < len(word2) {
			res = append(res, rune(word2[right]))
			right++
		}
	}

	return string(res)
}
