package leetcode

import (
	"math"
)

func LeftRightDifference(nums []int) []int {
	leftSum, rightSum, answer := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			leftSum[i], rightSum[len(nums)-1] = 0, 0
		} else {
			leftSum[i] = leftSum[i-1] + nums[i-1]
			rightSum[len(nums)-1-i] = rightSum[len(nums)-i] + nums[len(nums)-i]
		}
	}
	for i := 0; i < len(nums); i++ {
		answer[i] = int(math.Abs(float64(leftSum[i]) - float64(rightSum[i])))
	}
	return answer
}
