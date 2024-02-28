package leetcode

func MaxVowels(s string, k int) int {

	maxVowels := 0

	for i := 0; i < k; i++ {
		if vowelLetters(string(s[i])) {
			maxVowels++
		}
	}
	m := maxVowels
	l := k
	for l < len(s) {
		if vowelLetters(string(s[l-k])) {
			maxVowels--
		}
		if vowelLetters(string(s[l])) {
			maxVowels++
		}
		if maxVowels > m {
			m = maxVowels
		}
		l++
	}
	return m
}
func vowelLetters(s string) bool {
	if s == "a" || s == "u" || s == "i" || s == "e" || s == "o" {
		return true
	}
	return false
}
