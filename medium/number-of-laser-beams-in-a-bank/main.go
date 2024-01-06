package main

import "fmt"

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}), 8)
	fmt.Println(numberOfBeams([]string{"000", "111", "000"}), 0)
}

func numberOfBeams(bank []string) int {
	total := 0
	prevCnt := 0

	for _, row := range bank {
		var cnt int
		for _, b := range row {
			if string(b) == "1" {
				cnt++
			}
		}

		total += cnt * prevCnt
		if cnt > 0 {
			prevCnt = cnt
		}
	}
	return total
}
