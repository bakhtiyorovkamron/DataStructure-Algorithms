package leetcode

import (
	"strings"
)

func TruncateSentence(s string, k int) string {

	return strings.Join(strings.Split(s, " ")[0:k], " ")
}
