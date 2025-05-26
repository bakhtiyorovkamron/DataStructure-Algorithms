package leetcode

func LengthOfLastWord(s string) int {
	wordFound := false
	var newStr string
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			wordFound = true
		}
		if wordFound && string(s[i]) == " " {
			break
		}
		if wordFound {
			newStr += string(s[i])
		}
	}
	return len(newStr)
}
