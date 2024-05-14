package leetcode

func DifferenceOfSum(nums []int) int {
	res := 0
	for _, v := range nums {
		res += Abs(v - separate(v))
	}
	return res
}
func separate(a int) int {
	sum := 0
	for a >= 10 {
		sum += (a - (a/10)*10)
		a /= 10
	}
	return sum + a
}
