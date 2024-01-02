package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(decodeString("3[a]2[bc]"), "aaabcbc")
	fmt.Println(decodeString("3[a2[c]]"), "accaccacc")
	fmt.Println(decodeString("2[abc]3[cd]ef"), "abcabccdcdcdef")
}

func decodeString(s string) string {
	stack := make([]rune, 0, len(s))

	for _, w := range s {
		if string(w) != "]" {
			stack = append(stack, w)
			continue
		}

		decodedStr := make([]rune, 0, len(s))

		for string(stack[len(stack)-1]) != "[" {
			decodedStr = append(decodedStr, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = stack[:len(stack)-1]

		k := 0
		base := 1
		for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
			kTemp, _ := strconv.Atoi(string(stack[len(stack)-1]))
			k += kTemp * base
			base *= 10
			stack = stack[:len(stack)-1]
		}

		for k > 0 {
			for i := len(decodedStr) - 1; i >= 0; i-- {
				stack = append(stack, decodedStr[i])
			}
			k--
		}
	}

	return string(stack)
}
