package leetcode

import "fmt"

func ProductExceptSelf(nums []int) []int {
	//1, 2, 3, 4
	length := len(nums)

	answer := make([]int, length)

	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	fmt.Println("answer :", answer)

	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		fmt.Println("answer :", answer)
		R *= nums[i]
	}

	return answer
}
