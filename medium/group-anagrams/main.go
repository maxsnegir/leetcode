package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	cache := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		str := string(runes)
		if slice, ok := cache[str]; ok == true {
			slice = append(slice, s)
			cache[str] = slice
		} else {
			cache[str] = []string{s}
		}
	}

	for _, v := range cache {
		result = append(result, v)
	}
	return result
}
