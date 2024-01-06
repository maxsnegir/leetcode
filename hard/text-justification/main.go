package main

import (
	"fmt"
)

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

func fullJustify(words []string, maxWidth int) []string {
	buffer := make([]string, 0, maxWidth)
	totalLen := 0

	for _, word := range words {
		if len(word) < len(buffer) {
			buffer = append(buffer, word)
			totalLen += len(word)
			continue
		}

		buffer = make([]string, 0, maxWidth)
		buffer = append(buffer, word)
		totalLen = 0
	}

	return []string{}
}
