package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, word := range strs {
		if len(word) < len(prefix) {
			prefix = word
		}
	}

	for _, word := range strs {

		for prefix != "" {
			if word[:len(prefix)] == prefix {
				break
			}

			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return prefix
		}
	}

	return prefix
}
