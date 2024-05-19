package leetcode

import (
	"sort"
)

func MinOperations(nums []int, k int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] >= k {
			return i
		}
	}
	return 0
}
