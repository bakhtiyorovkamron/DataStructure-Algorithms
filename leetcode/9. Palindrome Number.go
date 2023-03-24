package leetcode

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	iks := fmt.Sprintf("%d", x)
	counter := len(iks) - 1
	for i := 0; i <= len(iks)/2; i++ {
		if string(iks[i]) != string(iks[counter]) {
			return false
		}
		counter--
	}
	return true
}
