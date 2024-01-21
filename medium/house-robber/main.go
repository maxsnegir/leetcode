package main

import (
	"fmt"
)

func main() {
	var res int

	res = rob([]int{1, 2, 3, 1})
	fmt.Println(res, "4")

	res = rob([]int{2, 7, 9, 3, 1})
	fmt.Println(res, "12")

	res = rob([]int{2, 1, 1, 2})
	fmt.Println(res, "4")
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1, len(nums)+1)
	dp[len(nums)] = 0
	dp[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		prev := dp[i+1]
		prevPlusCur := dp[i+2] + nums[i]
		dp[i] = max(prev, prevPlusCur)
	}

	return dp[0]
}
