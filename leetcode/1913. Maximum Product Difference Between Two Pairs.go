package leetcode

func MaxProductDifference(nums []int) int {
	nums = quickSortStart(nums)

	return Abs(nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1])
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
