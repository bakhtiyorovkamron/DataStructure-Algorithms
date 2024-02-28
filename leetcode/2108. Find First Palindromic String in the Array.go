package leetcode

func firstPalindrome(words []string) string {
	var isPalindrome = func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
