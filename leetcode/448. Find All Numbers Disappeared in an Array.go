package leetcode

func FindDisappearedNumbers(nums []int) []int {
	sort := make(map[int]struct{})
	l := len(nums)
	result := []int{}
	for _, v := range nums {
		if _, ok := sort[v]; !ok {
			sort[v] = struct{}{}
		}
	}
	for i := 1; i <= l; i++ {
		if _, ok := sort[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}
