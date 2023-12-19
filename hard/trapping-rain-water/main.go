package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
	fmt.Println(math.Ceil(float64(4) / float64(2)))

}

func trap(height []int) int {
	var res int
	left := 0
	right := len(height) - 1

	leftMax := 0
	rightMax := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}

	return res

}
