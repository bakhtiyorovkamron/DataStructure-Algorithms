package leetcode

func MoveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] == 0 {
				nums[j] = nums[j+1]
				nums[j+1] = 0
			}
		}
	}
}
