package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {

	arr := make([][]string, 0, numRows)

	for i, b := range s {

		arr[i%numRows] = append(arr[i%numRows], string(b))
	}

	return ""

}
