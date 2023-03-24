package leetcode

func SubtractProductAndSum(n int) int {
	forM, forA := 1, 0
	for n != 0 {
		forM *= (n % 10)
		forA += (n % 10)
		n /= 10
	}
	return forM - forA
}
