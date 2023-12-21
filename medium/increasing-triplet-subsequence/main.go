package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}), true)
	//fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}), false)
	//fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}), true)
	//fmt.Println(increasingTriplet([]int{2, 4, -2, -3}), false)
	fmt.Println(increasingTriplet([]int{1, 5, 0, 4, 1, 3}), true)
}

func increasingTriplet(nums []int) bool {

	firstNum := math.Inf(1)
	secondNum := math.Inf(1)

	for i := range nums {
		n := float64(nums[i])

		if n <= firstNum {
			firstNum = n
			continue
		}
		if n <= secondNum {
			secondNum = n
			continue
		}
		return true

	}
	return false
}

func increasingTripletTimeLimit(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}
