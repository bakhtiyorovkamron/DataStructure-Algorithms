package leetcode

func PlusOne(digits []int) []int {
	var sum = 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += sum
		if digits[i] > 9 {
			sum = digits[i] - 9
			digits[i] = 0
		} else {
			sum = 0
		}

	}
	if sum != 0 {
		newArr := []int{sum}
		newArr = append(newArr, digits...)
		return newArr
	}
	return digits
}
