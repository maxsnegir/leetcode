package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
}

func maxSubArray(nums []int) int {

	dp := make([]float64, len(nums))
	dp[0] = float64(nums[0])
	res := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = math.Max(dp[i-1]+float64(nums[i]), float64(nums[i]))
		res = math.Max(dp[i], res)
	}

	return int(res)
}
