package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeStars("leet**cod*e"), "lecoe")
	fmt.Println(removeStars("erase*****"), "")
}

func removeStars(s string) string {
	stack := make([]string, 0, len(s))

	for _, l := range s {

		if string(l) == "*" {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, string(l))
	}

	return strings.Join(stack, "")
}
