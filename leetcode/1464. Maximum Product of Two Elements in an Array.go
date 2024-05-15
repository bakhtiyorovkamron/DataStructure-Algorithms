package leetcode

func MaxProduct(nums []int) int {
	var (
		theIndex, firstBig, secondBig int
	)
	firstBig = nums[0]
	for i, v := range nums {
		if firstBig < v {
			firstBig, theIndex = v, i
		}
	}
	for i, v := range nums {
		if secondBig < v && i != theIndex {
			secondBig = v
		}
	}
	return (firstBig - 1) * (secondBig - 1)

}
