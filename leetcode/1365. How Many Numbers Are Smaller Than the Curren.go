package leetcode

import (
	"fmt"
	"sort"
)

func SmallerNumbersThanCurrent(nums []int) []int {

	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	m := make(map[int]int)
	fmt.Println("temp : ", temp)
	for i, v := range temp {
		if _, ok := m[v]; ok {
			continue
		} else {
			m[v] = i
		}
	}
	fmt.Println("nums : ", nums)
	for i, v := range nums {
		temp[i] = m[v]
	}
	fmt.Println("M : ", m)
	fmt.Println("temp : ", temp)
	return nums
}
