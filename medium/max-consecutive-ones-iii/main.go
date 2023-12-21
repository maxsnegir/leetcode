package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2), 6)
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3), 10)
	fmt.Println(longestOnes([]int{0, 0, 0, 1}, 4), 4)
}

func longestOnes(nums []int, k int) int {
	left := 0
	zeroCnt := k

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

	return len(nums) - left
}
