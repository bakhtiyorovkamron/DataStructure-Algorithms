package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
func lengthOfLastWord(s string) int {
	str := strings.Split(s, " ")
	for i := len(str) - 1; i >= 0; i-- {
		if len(str[i]) != 0 {
			return len(str[i])
		}
	}
	return -1
}
