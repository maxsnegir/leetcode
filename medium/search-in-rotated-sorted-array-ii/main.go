package main

import (
	"fmt"
)

func main() {
	//fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
	//fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13))
	//fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13))
	fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))
}

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return true
		}

		if nums[left] == nums[middle] {
			left++
			continue
		}

	}

	return false

}

func bSearch(nums []int, l, r, target int) bool {
	left := l
	right := r

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
