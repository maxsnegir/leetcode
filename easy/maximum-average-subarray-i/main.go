package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4), 12.75000)
	fmt.Println(findMaxAverage([]int{5}, 1), 5.00)
	fmt.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1), 4)
}

func findMaxAverage(nums []int, k int) float64 {

	var max float64

	for _, v := range nums[:k] {
		max += float64(v)
	}

	tmpMax := max
	for i := k; i < len(nums); i++ {
		tmpMax += float64(nums[i] - nums[i-k])
		max = math.Max(tmpMax, max)
	}

	return max / float64(k)
}
