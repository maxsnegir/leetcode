package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(predictTheWinner([]int{1, 5, 2}), false)
	fmt.Println(predictTheWinner([]int{1, 5, 233, 7}), true)
}

func predictTheWinner(nums []int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	steps := int(math.Ceil(float64(len(nums)) / 2))
	r1 := helper(nums[1:], nums[0], false, steps-1)
	r2 := helper(nums[:len(nums)-1], nums[len(nums)-1], false, steps-1)

	return r1 >= s-r1 || r2 > s-r2
}

func helper(nums []int, s int, skip bool, steps int) int {
	if steps == 0 {
		return s
	}

	s1 := s
	s2 := s

	if !skip {
		s1 += nums[0]
		s2 += nums[len(nums)-1]
	}

	a := helper(nums[1:], s1, !skip, steps-1)
	b := helper(nums[:len(nums)-1], s2, !skip, steps-1)
	return int(math.Max(float64(a), float64(b)))
}
