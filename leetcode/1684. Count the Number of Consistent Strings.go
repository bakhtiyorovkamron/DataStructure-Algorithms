package leetcode

import (
	"strings"
)

func CountConsistentStrings(allowed string, words []string) int {
	counter := 0
	for _, v := range words {
		isAllowed := true
		for _, c := range v {
			if !strings.Contains(allowed, string(c)) {
				isAllowed = false
			}
		}
		if isAllowed {
			counter++
		}
	}
	return counter
}
