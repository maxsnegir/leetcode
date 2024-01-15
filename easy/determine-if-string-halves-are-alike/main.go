package main

import (
	"fmt"
)

func main() {
	fmt.Println(halvesAreAlike("book"), true)
	fmt.Println(halvesAreAlike("textbook"), false)
	fmt.Println(halvesAreAlike("AbCdEfGh"), true)
}

func halvesAreAlike(s string) bool {
	vowels := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
		"A": {},
		"E": {},
		"I": {},
		"O": {},
		"U": {},
	}

	left := s[:len(s)/2]
	right := s[len(s)/2:]

	var res int
	for i := 0; i < len(left); i++ {
		_, ok := vowels[string(left[i])]
		if ok {
			res++
		}
	}

	for i := 0; i < len(right); i++ {
		_, ok := vowels[string(right[i])]
		if ok {
			res--
		}
	}

	return res == 0
}
