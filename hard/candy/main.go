package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(candy([]int{1, 0, 2}), 5)
	fmt.Println(candy([]int{1, 3, 2, 2, 1}), 7)
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}), 13)
}

func candy(ratings []int) int {

	dp := make([]int, len(ratings))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}

	s := dp[len(dp)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			dp[i] = int(math.Max(float64(dp[i]), float64(dp[i+1]+1)))
		}

		s += dp[i]
	}

	return s

}
