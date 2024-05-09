package leetcode

func Shuffle(nums []int, n int) []int {
	newarr := make([]int, n*2)
	index := 0
	for i := 0; i < n; i++ {
		newarr[index], newarr[index+1] = nums[i], nums[i+n]
		index += 2
	}
	return newarr
}
