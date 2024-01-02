package main

import (
	"fmt"
)

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGain := 0
	currGain := 0
	answer := 0
	for i := 0; i < len(gas); i++ {
		totalGain += gas[i] - cost[i]
		currGain += gas[i] - cost[i]

		if currGain < 0 {
			currGain = 0
			answer += 1
		}
	}

	if totalGain < 0 {
		return -1
	}
	return gas[answer]
}
