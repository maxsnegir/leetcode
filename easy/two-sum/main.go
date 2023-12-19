package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9), "[0,1]")
	fmt.Println(twoSum([]int{3, 2, 4}, 6), "[1,2]")
	fmt.Println(twoSum([]int{3, 3}, 6), "[0, 1]")

}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i, num := range nums {

		want := target - num
		if idx, ok := cache[want]; ok == true {
			return []int{idx, i}
		}

		cache[num] = i

	}

	return []int{}
}
