package leetcode

func BuildArray(nums []int) []int {
	// 0,2,1,5,3,4
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
