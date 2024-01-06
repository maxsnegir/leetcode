package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minOperations([]int{2, 3, 3, 2, 2, 4, 2, 3, 4}), 4)
	fmt.Println(minOperations([]int{2, 1, 2, 2, 3, 3}), -1)
	fmt.Println(minOperations([]int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}), 7)
}

func minOperations(nums []int) int {
	counter := make(map[int]int)

	for _, val := range nums {
		counter[val]++
	}

	var res float64
	for _, v := range counter {
		if v == 1 {
			return -1
		}

		res += math.Ceil(float64(v) / 3)
	}

	return int(res)
}
