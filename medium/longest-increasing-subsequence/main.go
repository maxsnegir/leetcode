package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 2, 5, 3, 7, 101, 101, 18}), 4)
	fmt.Println(lengthOfLIS([]int{7, 7, 7}), 1)
}

func lengthOfLIS(nums []int) int {
	return 0
}
