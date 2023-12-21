package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/string-compression/

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'a'}), 1, "a3")
	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b'}), 3, "a3b2")
}

func compress(chars []byte) int {
	curChar := chars[0]
	writeIdx := 1
	cnt := 1

	for i := 1; i < len(chars); i++ {
		if chars[i] == curChar {
			cnt++
			continue
		}

		if cnt != 1 {
			digits := []byte(strconv.Itoa(cnt))
			for _, d := range digits {
				chars[writeIdx] = d
				writeIdx++
			}
		}

		chars[writeIdx] = chars[i]
		curChar = chars[writeIdx]
		writeIdx++
		cnt = 1
	}

	if cnt != 1 {
		digits := []byte(strconv.Itoa(cnt))
		for _, d := range digits {
			chars[writeIdx] = d
			writeIdx++
		}
	}

	return writeIdx
}
