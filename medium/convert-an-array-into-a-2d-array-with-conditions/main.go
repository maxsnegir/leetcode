package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}), "[[1,3,4,2],[1,3],[1]]")
	fmt.Println(findMatrix([]int{1, 2, 3, 4}), "[[4,3,2,1]]")
}

func findMatrix(nums []int) [][]int {

	freq := make(map[int]int)

	for _, val := range nums {
		freq[val]++
	}

	result := make([][]int, 0, len(nums))

	for k, v := range freq {

		for i := 1; i <= v; i++ {
			if i > len(result) {
				result = append(result, []int{})
			}

			result[i-1] = append(result[i-1], k)
		}
	}

	return result
}
