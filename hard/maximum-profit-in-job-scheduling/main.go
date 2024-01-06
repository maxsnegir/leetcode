package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	fmt.Println(jobScheduling([]int{1, 2, 2, 3}, []int{2, 5, 3, 4}, []int{3, 4, 1, 2}))
	fmt.Println(jobScheduling([]int{4, 2, 4, 8, 2}, []int{5, 5, 5, 10, 8}, []int{1, 2, 8, 10, 4}))
	fmt.Println(jobScheduling([]int{6, 15, 7, 11, 1, 3, 16, 2}, []int{19, 18, 19, 16, 10, 8, 19, 8}, []int{2, 9, 1, 19, 5, 7, 3, 19}))
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	intervals := make([][]int, len(startTime))
	for i := range startTime {
		intervals[i] = []int{startTime[i], endTime[i], profit[i]}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] && intervals[i][1] == intervals[j][1] {
			return intervals[i][2] <= intervals[j][2]
		}
		return intervals[i][0] <= intervals[j][0]
	})

	dp := make([]int, 0, len(intervals))
	dp = append(dp, intervals[0][2])
	maximum := 0

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		dp = append(dp, intervals[i][2])
		tmpRes := intervals[i][2]

		for j := i - 1; j >= 0; j-- {

			prev := intervals[j]

			if cur[0] >= prev[1] && cur[1] > prev[0] {
				tmpRes = max(tmpRes, dp[i]+dp[j])
			}
		}
		dp[i] = tmpRes
		maximum = max(maximum, tmpRes)

	}
	return maximum
}
