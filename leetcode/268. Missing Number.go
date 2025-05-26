package leetcode

func MissingNumber(nums []int) int {
	max := nums[0]
	min := nums[0]
	elemInOrders := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		elemInOrders[nums[i]] = struct{}{}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	if min != 0 {
		return 0
	}
	for i := min; i <= max; i++ {
		if _, ok := elemInOrders[i]; !ok {
			return i
		}
	}
	return max + 1
}
