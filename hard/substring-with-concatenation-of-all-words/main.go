package main

import (
	"fmt"
)

// ToDo

func main() {
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}), "[0, 9]")
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}), "[]")
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}), "[6, 9, 12]")
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0, len(s))

	lenWord := len(words[0])
	needLenWord := len(words) * lenWord

	wordsSet := make(map[string]struct{})
	for _, word := range words {
		wordsSet[word] = struct{}{}
	}

	left := 0
	for i := 0; i < len(s); i += lenWord {
		word := s[i : i+lenWord]

		if _, ok := wordsSet[word]; !ok {
			left = i
			continue
		}

		if len(s[left:i]) == needLenWord {
			res = append(res, left)
			left += lenWord
		}

	}

	return res
}
