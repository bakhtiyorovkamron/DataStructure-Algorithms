package leetcode

func ContainsDuplicate(nums []int) bool {
	numbers := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numbers[nums[i]]; ok {
			return true
		} else {
			numbers[nums[i]] = i
		}
	}
	return false
}
