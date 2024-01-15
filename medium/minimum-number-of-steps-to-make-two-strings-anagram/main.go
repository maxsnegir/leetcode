package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSteps("bab", "aba"), 1)
	fmt.Println(minSteps("leetcode", "practice"), 5)
}

func minSteps(s string, t string) int {
	cnt := make(map[string]int)

	for i := range s {
		cnt[string(s[i])]++
		cnt[string(t[i])]--
	}

	var res int

	for _, vs := range cnt {
		res += max(0, vs)
	}

	return res
}
