package leetcode

func MajorityElement(nums []int) int {
	l := len(nums)
	elements := make(map[int]int)
	if l == 1 {
		return nums[0]
	}
	theMajorityNum := 0
	theIndex := 0

	for i := 0; i < len(nums); i++ {
		if _, ok := elements[nums[i]]; ok {
			elements[nums[i]]++
			if elements[nums[i]] > (l / 2) {
				if elements[nums[i]] > theIndex {
					theMajorityNum = nums[i]
					theIndex = elements[nums[i]]
				}
			}
		} else {
			elements[nums[i]] = 1
			if len(elements) == 1 {
				theMajorityNum = nums[i]
				theIndex = 1
			}
		}
	}

	return theMajorityNum
}
