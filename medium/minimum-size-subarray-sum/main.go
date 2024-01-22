package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), "2")
	//fmt.Println(minSubArrayLen(4, []int{1, 4, 4}), "1")
	//fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}), "0")
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}), "3")
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := math.Inf(1)
	s := 0
	left := 0

	for i := 0; i < len(nums); i++ {
		s += nums[i]

		for s >= target {
			res = min(res, float64(i-left+1))
			s -= nums[left]
			left++
		}
	}

	if res == math.Inf(1) {
		return 0
	}
	return int(res)
}
