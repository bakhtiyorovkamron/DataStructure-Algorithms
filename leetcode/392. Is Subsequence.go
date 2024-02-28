package leetcode

func IsSubsequence(s string, t string) bool {

	if s == "" {
		return false
	}
	sIndex := 0

	for tIndex := range t {
		if len(s) > sIndex && s[sIndex] == t[tIndex] {
			sIndex++
		}
	}
	return sIndex == len(s)
}
