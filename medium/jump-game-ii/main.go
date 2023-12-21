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

	prevMin := math.Inf(1)

	for position := len(nums) - 2; position >= 0; position-- {
		jumps := nums[position]
		if position+jumps >= len(nums)-1 {
			dp[position] = 1
			prevMin = 1
			continue
		}
		dp[position] = math.Inf(1)
		if jumps == 0 {
			continue
		}

		for j := position + 1; j <= position+jumps; j++ {
			m := math.Min(dp[j], dp[position])
			if m == prevMin {
				dp[position] = prevMin
				break
			}
			dp[position] = m
		}
		dp[position]++
	}

	return int(dp[0])
}
