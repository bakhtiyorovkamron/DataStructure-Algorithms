package leetcode

func AreAlmostEqual(s1 string, s2 string) bool {
	indexes := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			indexes = append(indexes, i)
		}
	}
	if s1[indexes[0]] == s2[indexes[1]] && s1[indexes[1]] == s2[indexes[0]] {
		return true
	}
	return false
}
