package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(7))
	fmt.Println(isHappy(116))
}

func process(n int) int {
	var digits []int
	for n/10 != 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	s := n * n

	for _, d := range digits {
		s += d * d
	}
	return s
}

func isHappy(n int) bool {
	cache := make(map[int]struct{})

	for {
		if n == 1 {
			return true
		}

		n = process(n)
		if _, ok := cache[n]; ok {
			return false
		}
		cache[n] = struct{}{}
	}

}
