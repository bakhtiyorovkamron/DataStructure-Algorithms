package leetcode

func CountKDifference(nums []int, k int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for ; j < len(nums); j++ {
			if Abs(nums[i]-nums[j]) == k {
				counter++
			}
		}
	}
	return counter
}
