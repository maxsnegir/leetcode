package main

import (
	"fmt"
)

func main() {

	fmt.Println(judgeCircle("UD"), true)
	fmt.Println(judgeCircle("LL"), false)
	fmt.Println(judgeCircle("LR"), true)
	fmt.Println(judgeCircle("UR"), false)
}

func judgeCircle(moves string) bool {
	expected := [2]int{0, 0}
	start := [2]int{0, 0}

	for _, r := range moves {
		move := string(r)

		switch move {
		case "R":
			start[1] += 1
		case "L":
			start[1] -= 1
		case "U":
			start[0] += 1
		case "D":
			start[0] -= 1
		}

	}
	return start == expected
}
