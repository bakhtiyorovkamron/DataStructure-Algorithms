package leetcode

import (
	"math"
)

func IncreasingTriplet(nums []int) bool {
	//1,5,0,4,1,3
	triplet1, triplet2 := math.MaxInt32, math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] <= triplet1 {
			triplet1 = nums[i]
		} else if nums[i] <= triplet2 {
			triplet2 = nums[i]
		} else {
			return true
		}
	}
	return false
}
