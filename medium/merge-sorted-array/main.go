package main

import (
	"fmt"
	"sort"
)

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	fmt.Println(n1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, val := range nums2 {
		nums1[m+i] = val
	}

	sort.Ints(nums1)
}
