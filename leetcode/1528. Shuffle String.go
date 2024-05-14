package leetcode

import "sort"

func restoreString(s string, indices []int) string {
	shuffled := make(map[int]string)
	var result string
	for i := range indices {
		shuffled[indices[i]] = string(s[i])
	}
	sort.Ints(indices)
	for i := range indices {
		result += shuffled[indices[i]]
	}
	return result
}
