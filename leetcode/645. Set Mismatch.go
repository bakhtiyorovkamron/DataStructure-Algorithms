package leetcode

import "fmt"

func FindErrorNums(nums []int) []int {
	n := 1
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i], n)
		if nums[i] != n {
			return []int{nums[i], n}
		}
		n++
	}
	return []int{}
}
