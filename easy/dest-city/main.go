package main

import (
	"fmt"
)

func main() {
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
}

func destCity(paths [][]string) string {
	cache := make(map[string]string)

	for _, path := range paths {
		arr := path[0]  // Z
		dest := path[1] // B

		cache[arr] = dest
		if _, ok := cache[dest]; ok != true {
			cache[dest] = ""
		}
	}

	for k, v := range cache {
		if v == "" {
			return k
		}
	}
	return ""
}
