package main

import (
	"fmt"
)

func main() {
	//fmt.Println(isIsomorphic("egg", "att"), true)
	//fmt.Println(isIsomorphic("foo", "bar"), false)
	//fmt.Println(isIsomorphic("paper", "title"), true)
	fmt.Println(isIsomorphic("badc", "baba"), false)
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	replacedL := make(map[string]string)

	for i := 0; i < len(s); i++ {
		l := string(s[i])
		t := string(t[i])
		if newL, ok := replacedL[l]; ok {
			if newL != t {
				return false
			}
			continue
		}

		if l != t {
			replacedL[l] = t
		} else {
			replacedL[l] = l
		}
	}
	return true
}
