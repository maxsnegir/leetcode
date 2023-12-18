package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/container-with-most-water/submissions/
func main() {

	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	fmt.Println(maxArea([]int{1, 1}), 1)

}
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var area float64

	for left <= right {

		distance := right - left

		leftH := height[left]
		rightH := height[right]
		LArea := float64(leftH * distance)
		RArea := float64(rightH * distance)

		area = math.Max(area, math.Min(LArea, RArea))
		if leftH > rightH {
			right--
		} else {
			left++
		}

	}

	return int(area)
}
