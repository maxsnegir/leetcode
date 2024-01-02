package main

import (
	"fmt"
)

func main() {
	n := []int{3, 2, 2, 3}
	fmt.Println(removeElement(n, 3), "|", "[2, 2]", "|", n)
	n = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(n, 2), "|", "[0,1,4,0,3]", "|", n)
}

func removeElement(nums []int, val int) int {
	left := 0

	for i, num := range nums {
		if num != val {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			continue
		}

	}

	return left
}
