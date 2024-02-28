package leetcode

func ClimbStairs(n int) int {
	j := 1
	temp := j
	for i := 1; i < n; i++ {
		j = j + temp
		temp = j - temp
	}
	return j
}
