package leetcode

func SingleNumber2(nums []int) int {
	n := make(map[int]int)
	for _, v := range nums {
		n[v]++
	}
	for k, v := range n {
		if v == 1 {
			return k
		}
	}
	return 0
}
