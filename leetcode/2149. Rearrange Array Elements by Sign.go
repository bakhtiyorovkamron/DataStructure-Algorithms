package leetcode

func RearrangeArray(nums []int) []int {
	negative, positive := []int{}, []int{}
	for _, num := range nums {
		if num < 0 {
			negative = append(negative, num)
		} else {
			positive = append(positive, num)
		}
	}
	finallist := []int{}
	for i := 0; i < len(positive); i++ {
		finallist = append(finallist, positive[i])
		if i < len(negative) {
			finallist = append(finallist, negative[i])
		}
	}
	return finallist
}
