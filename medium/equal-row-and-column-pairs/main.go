package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}

func equalPairs(grid [][]int) int {
	var res int
	rowsMap := make(map[string]int)

	for _, row := range grid {
		s := make([]rune, 0, len(grid))

		for _, b := range row {
			s = append(s, rune(b))
		}
		rowsMap[string(s)]++
	}

	for column := 0; column < len(grid); column++ {
		s := make([]rune, 0, len(grid))

		for row := 0; row < len(grid); row++ {
			s = append(s, rune(grid[row][column]))
		}

		res += rowsMap[string(s)]
	}

	return res

}
