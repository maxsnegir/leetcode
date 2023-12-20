package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"), "ABC")
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"), "AB")
	fmt.Println(gcdOfStrings("LEET", "CODE"), "")
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"), "TAUXX")
	fmt.Println(gcdOfStrings("ABCABCABC", "ABCAAA"), "")
}

func gcdOfStrings(largest string, smaller string) string {

	if largest < smaller {
		largest, smaller = smaller, largest
	}

	i := len(smaller)

	for ; i > 0; i-- {
		if len(largest)%i != 0 || len(smaller)%i != 0 {
			continue
		}

		if check(largest, smaller, i) {
			break
		}
	}

	return smaller[:i]
}

func check(largest, smaller string, i int) bool {
	for j := 0; j < len(largest); j += i {
		if largest[j:j+i] != smaller[:i] {
			return false
		}
	}
	for j := 0; j < len(smaller); j += i {
		if smaller[j:j+i] != smaller[:i] {
			return false
		}
	}
	return true
}
