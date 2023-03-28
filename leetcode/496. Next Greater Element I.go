package leetcode

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	result := []int{}
	for i := 0; i < len(nums1); i++ {
		indexInNums2 := search(nums2, nums1[i])
		n := func(index, target int, arr []int) int {
			for ; index < len(arr); index++ {
				if arr[index] > target {
					return arr[index]
				}
			}
			return -1
		}
		result = append(result, n(indexInNums2, nums1[i], nums2))
	}
	return result
}
func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
