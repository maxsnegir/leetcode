package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{7, 0, 1, 2}, 2))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
}

func search(nums []int, target int) int {

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

	res := bSearch(nums, 0, left-1, target)
	if res != -1 {
		return res
	}
	return bSearch(nums, left, len(nums)-1, target)
}

func bSearch(nums []int, l, r, target int) int {
	left := l
	right := r

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right--
		} else {
			left++
		}
	}
	return -1
}
