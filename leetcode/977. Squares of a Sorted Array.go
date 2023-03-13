package leetcode

import "fmt"

func SortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}
func Rotate(nums []int, k int) {
	k = k % len(nums)
	if k != 0 {
		fmt.Println("append :",append(nums[len(nums)-k:], nums[:len(nums)-k]...))
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	}
}
