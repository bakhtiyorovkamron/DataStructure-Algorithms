package leetcode

import "fmt"

//3,2,2,3
func RemoveElement(nums []int, val int) int {
	index := 0
	// if nums[0] == val {
	// 	index = 0 
	// }
	for i := 0 ; i < len(nums) ; i++ {
		if nums[i] == val {
			index = i
		}else{
			nums[index] = nums[i]
			// nums[i] = nums[index]
			// index = i
		}
	}
	fmt.Println(nums)
	return 9
}