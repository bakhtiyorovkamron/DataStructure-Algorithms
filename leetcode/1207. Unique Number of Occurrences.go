package leetcode

func UniqueOccurrences(arr []int) bool {

	numbers := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		numbers[arr[i]]++
	}
	sum := 0
	for _, v := range numbers {
		for _, w := range numbers {
			if w == v {
				sum++
			}
		}
		if sum > 1 {
			return false
		}
		sum = 0
	}
	return true
}
