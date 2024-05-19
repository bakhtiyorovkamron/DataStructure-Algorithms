package leetcode

func DecompressRLElist(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums)/2; i++ {
		freq, val := nums[2*i], nums[2*i+1]
		for j := 0; j < freq; j++ {
			result = append(result, val)
		}
	}
	return result
}
