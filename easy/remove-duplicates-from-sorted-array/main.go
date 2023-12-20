package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3}), 3)
}

func removeDuplicates(nums []int) int {

	insertIndx := 1

	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		prev := nums[i-1]

		if prev == cur {
			continue
		}

		nums[insertIndx] = cur
		insertIndx++
	}

	return insertIndx
}
