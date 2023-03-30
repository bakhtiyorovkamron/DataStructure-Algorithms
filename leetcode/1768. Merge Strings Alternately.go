package leetcode

func MergeAlternately(word1 string, word2 string) string {
	mergedWord := ""
	l := len(word1)
	if len(word1) > len(word2) {
		l = len(word2)
	} else if len(word1) < len(word2) {
		l = len(word1)
	}
	for i := 0; i < len(word1); i++ {
		mergedWord += string(word1[i]) + string(word2[i])
		if i == l-1 {
			i++
			if len(word1) > len(word2) {
				for ; i < len(word1); i++ {
					mergedWord += string(word1[i])
				}
			} else if len(word1) < len(word2) {
				for ; i < len(word2); i++ {
					mergedWord += string(word2[i])
				}
			}
		}
	}
	return mergedWord
}
