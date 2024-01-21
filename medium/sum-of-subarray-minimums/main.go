package main

import (
	"fmt"
	"math"
)

func main() {
	var res int

	res = sumSubarrayMins([]int{3, 1, 2, 4})
	fmt.Println(res, "17")

	res = sumSubarrayMins([]int{11, 81, 94, 43, 3})
	fmt.Println(res, "444")
}

func sumSubarrayMins(arr []int) int {
	res := 0
	mod := int(math.Pow(10, 9) + 7)

	for _, v := range arr {
		res += v
	}

	for i := 0; i < len(arr); i++ {
		m := arr[i]
		for j := i + 1; j < len(arr); j++ {
			m = min(m, arr[j])
			res += m
		}
	}

	return res % mod
}
