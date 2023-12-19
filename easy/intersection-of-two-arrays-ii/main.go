package main

import (
	"fmt"
)

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	counters := make(map[int]int)

	for _, num := range nums1 {
		if _, ok := counters[num]; ok {
			counters[num]++
		} else {
			counters[num] = 1
		}
	}

	res := make([]int, 0, len(nums2))

	for _, num := range nums2 {
		if val, ok := counters[num]; ok && val > 0 {
			res = append(res, num)
			counters[num]--
		}
	}

	return res

}
