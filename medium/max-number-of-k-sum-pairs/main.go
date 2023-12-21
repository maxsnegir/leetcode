package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxOperations([]int{1, 1, 2, 2, 4, 6}, 5), 1)
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5), 2)
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	cnt := 0

	for left < right {
		s := nums[left] + nums[right]

		if s == k {
			cnt++
			left++
			right--
			continue
		}

		if s > k {
			right--
			continue
		}
		left++
	}
	return cnt
}
