package main

import (
	"fmt"
)

func main() {

	fmt.Println(longestSubarray([]int{1, 1, 0, 1}), 3)
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5)
}

func longestSubarray(nums []int) int {
	left := 0
	zeroCnt := 1

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCnt--
		}

		if zeroCnt < 0 {
			if nums[left] == 0 {
				zeroCnt++
			}
			left++
		}
	}

	return len(nums) - left - 1
}
