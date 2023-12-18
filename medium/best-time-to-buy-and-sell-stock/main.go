package main

import (
	"fmt"
	"math"
)

// [7,1,5,3,6,4] -> 5

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}), 0)
}

func maxProfit(prices []int) int {
	var maxProf float64
	minPrice := math.Inf(1)

	for _, price := range prices {
		minPrice = math.Min(float64(price), minPrice)
		maxProf = math.Max(maxProf, float64(price)-minPrice)
	}

	return int(maxProf)
}
