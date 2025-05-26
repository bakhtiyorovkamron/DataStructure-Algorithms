package leetcode

// import "fmt"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 { return "" }
	var str string
	for  i := 0 ;i < len(strs[0]); i++ {
		firstLetter := (strs[0][i])
		count := 1
		for j := 1 ; j < len(strs) ; j++ {
			if len(strs[j]) > i && strs[j][i]==firstLetter {
				count++
			}
		}
		if count != len(strs) {
			return str
		}
		str += string(firstLetter)
	}
	return str
}
