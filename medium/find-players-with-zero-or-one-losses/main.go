package main

import (
	"fmt"
	"sort"
)

func main() {
	var res [][]int
	res = findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}})
	fmt.Println(res, "[[1,2,10],[4,5,7,8]]")
}

func findWinners(matches [][]int) [][]int {
	winners := make(map[int]int)
	lossers := make(map[int]int)

	winnersRes := make([]int, 0, len(winners))
	lossersRes := make([]int, 0, len(lossers))

	for _, match := range matches {
		w := match[0]
		l := match[1]

		if _, ok := lossers[w]; ok {
			delete(winners, w)
			lossers[l]++
			continue
		}

		winners[w]++
		lossers[l]++
	}

	for k, _ := range winners {
		if _, ok := lossers[k]; !ok {
			winnersRes = append(winnersRes, k)
		}

	}

	for k, v := range lossers {
		if v == 1 {
			lossersRes = append(lossersRes, k)
		}
	}

	sort.Ints(winnersRes)
	sort.Ints(lossersRes)
	return [][]int{winnersRes, lossersRes}
}
