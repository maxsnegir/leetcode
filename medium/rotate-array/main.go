package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(x, 3)
	fmt.Println(x, "[5,6,7,1,2,3,4]")
}

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1]

		for j := 0; j < len(nums); j++ {
			nums[j], prev = prev, nums[j]
			fmt.Println(nums)
		}
	}
}
