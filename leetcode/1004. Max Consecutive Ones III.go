package leetcode

func LongestOnes(nums []int, k int) int {

	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			max++
		}
		if nums[i] == 0 {
			max++
			k--
		}
	}

	return 0
}
