package leetcode

import (
	"strings"
)

func Interpret(command string) string {
	parsedString := ""
	for i := 0; i < len(command); i++ {
		if string(command[i]) == "(" && string(command[i+1]) == ")" {
			parsedString += "o"
			i++
		} else if string(command[i]) == "(" && string(command[i+1]) == "a" && string(command[i+2]) == "l" && string(command[i+3]) == ")" {
			parsedString += "al"
			i += 3
		} else {
			parsedString += string(command[i])
		}
	}
	return strings.ReplaceAll(strings.ReplaceAll(command, "()", "o"), "(al)", "al")
}
