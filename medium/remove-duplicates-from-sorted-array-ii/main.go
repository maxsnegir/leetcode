package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}), 5)
}

func removeDuplicates(nums []int) int {

	count := 1
	j := 1

	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[i] = nums[j]
			j++
		}

	}

	return j
}
