package leetcode

func FindMaxConsecutiveOnes(nums []int) int {

	max1 := 0
	max2 := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			max1++
			if max1 > max2 {
				max2 = max1
			}
		} else {
			max1 = 0
		}
	}
	return max2
}
