package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-the-highest-altitude

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}), 1)
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}), 0)
}

func largestAltitude(gain []int) int {

	s := 0
	res := 0
	for _, g := range gain {
		s += g
		if s > res {
			res = s
		}
	}
	return res
}
