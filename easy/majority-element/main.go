package main

import (
	"math"
)

func main() {

}
func majorityElement(nums []int) int {

	cnt := int(math.Ceil(float64(len(nums)) / 2))
	cache := make(map[int]int)

	for _, num := range nums {
		cache[num]++
	}

	for k, v := range cache {
		if v >= cnt {
			return k
		}
	}
	return 0
}
