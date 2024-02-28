package leetcode

func FindErrorNums(nums []int) []int {
	var dup, missing = 0,1
	for i := 0; i < len(nums); i++ {
		var count = 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count == 2 {
			dup = i
		} else {
			missing = i
		}
	}
	return []int{dup, missing}
}
