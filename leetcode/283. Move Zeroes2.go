package leetcode

import "fmt"

func MoveZeroe2s(nums []int) {
	//0, 1, 0, 3, 12
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	fmt.Println(nums)
}
