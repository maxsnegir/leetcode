package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords(" 1 2 3 "), "3 2 1")
	fmt.Println(reverseWords("the sky is blue"), "blue is sky the")
	fmt.Println(reverseWords("a good   example"), "example good a")
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	words := strings.Split(s, " ")

	stack := make([]string, 0, len(words))

	for _, w := range words {
		if string(w) == "" {
			continue
		}
		stack = append(stack, w)
	}

	r := make([]string, 0, len(words))

	for _ = range stack {
		r = append(r, stack[len(stack)-1])
		stack = stack[:len(stack)-1]

	}

	return strings.Join(r, " ")
}

//func reverseWords(s string) string {
//	s = strings.Trim(s, " ")
//
//	words := strings.Split(s, " ")
//
//	left := 0
//	right := len(words) - 1
//
//	var hasSpaces bool
//	for left < right {
//		leftWord := words[left]
//		rightWord := words[right]
//
//		if leftWord == "" {
//			left++
//			hasSpaces = true
//			continue
//		}
//		if rightWord == "" {
//			right--
//			hasSpaces = true
//			continue
//		}
//
//		words[left], words[right] = rightWord, leftWord
//		left++
//		right--
//	}
//
//	if hasSpaces {
//		newWords := make([]string, 0, len(words))
//		for _, w := range words {
//			if string(w) != "" {
//				newWords = append(newWords, w)
//			}
//		}
//		return strings.Join(newWords, " ")
//	}
//
//	return strings.Join(words, " ")
//}
