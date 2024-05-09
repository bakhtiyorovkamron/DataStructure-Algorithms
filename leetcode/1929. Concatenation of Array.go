package leetcode



func GetConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i+len(nums)] = nums[i]
	}
    return ans
}