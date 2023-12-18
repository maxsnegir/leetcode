package main

import (
	"fmt"
	"sort"
)

func main() {

	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-2, 0, 2}))
	//fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
			if sum > 0 {
				right--
			} else {
				left++
			}

		}
	}

	return result
}
