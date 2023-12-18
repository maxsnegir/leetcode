package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}), true)
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}), false)
	fmt.Println(canJump([]int{0}), true)
	fmt.Println(canJump([]int{2, 0, 0}), true)
	fmt.Println(canJump([]int{3, 0, 8, 2, 0, 0, 1}), true)
}

func canJump(nums []int) bool {

	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	end := len(dp) - 1

	for pos := len(nums) - 2; pos >= 0; pos-- {

		jumps := nums[pos]

		for jumps > 0 {
			if jumps+pos >= end {
				dp[pos] = true
				break
			}
			if dp[jumps+pos] == true {
				dp[pos] = true
			}
			jumps--
		}
	}

	return dp[0]
}
