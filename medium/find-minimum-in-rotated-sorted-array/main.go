package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{0, 1, 2, 3}))
	fmt.Println(findMin([]int{1, 2, 3, 0}))
	fmt.Println(findMin([]int{2, 4, 0, 1}))
}

func findMin(nums []int) int {
	left := 0
	right := len(nums)

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] > nums[len(nums)-1] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return nums[left]
}
