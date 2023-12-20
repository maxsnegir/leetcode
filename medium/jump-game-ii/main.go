package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}), 2)
	fmt.Println(jump([]int{2, 3, 0, 1, 4}), 2)
}

func jump(nums []int) int {
	dp := make([]float64, len(nums))

	for pos := len(nums) - 2; pos >= 0; pos-- {

		jumps := nums[pos]
		if pos+jumps >= len(nums)-1 {
			dp[pos] = 1
			continue
		}
		m := math.Inf(1)
		dp[pos] = m
		if jumps == 0 {
			continue
		}

		for j := pos + 1; j <= pos+jumps; j++ {
			dp[pos] = math.Min(dp[j], dp[pos])
		}
		dp[pos]++
	}

	return int(dp[0])
}
