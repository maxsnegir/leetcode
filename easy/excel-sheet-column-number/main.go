package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("A"), 1)
	fmt.Println(titleToNumber("B"), 2)
	fmt.Println(titleToNumber("Z"), 26)
	fmt.Println(titleToNumber("AA"), 27)
	fmt.Println(titleToNumber("AB"), 27)

}
func titleToNumber(columnTitle string) int {
	letters := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	columns := make(map[string]int)

	number := 1
	for _, letter := range letters {
		columns[letter] = number
		number++
	}

	var res int

	for i := 0; i < len(columnTitle); i++ {
		l := string(columnTitle[i])
		res += columns[l]
		res *= 26
	}

	return res / 26
}
