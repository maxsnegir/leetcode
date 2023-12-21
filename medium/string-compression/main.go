package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b', 'c'}))
}

func compress(chars []byte) int {

	left := 1
	right := left + 1

	for right < len(chars) {
		if chars[right] == chars[left] {
			right++
			continue
		}

		digits := []byte(strconv.Itoa(right - left))
		for i := range digits {
			chars[left] = digits[i]
			left++
		}

		chars[left] = chars[right]
		left = right
		right++
		fmt.Println(string(chars))
	}

	return left

}
