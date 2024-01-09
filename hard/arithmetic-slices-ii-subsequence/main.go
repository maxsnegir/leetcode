package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}), 7)
}

func numberOfArithmeticSlices(nums []int) int {
	ans := 0
	cnt := make([]map[int]int, len(nums))

	for i := range nums {
		cnt[i] = make(map[int]int)
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			delta := nums[i] - nums[j]
			s := 0

			if _, ok := cnt[j][delta]; ok {
				s = cnt[j][delta]
			}

			cnt[i][delta] += s + 1
			ans += s
		}
	}

	return ans
}
