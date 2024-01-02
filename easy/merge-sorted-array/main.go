package main

import (
	"sort"
)

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, val := range nums2 {
		nums1[m+i] = val
	}

	sort.Ints(nums1)
}
