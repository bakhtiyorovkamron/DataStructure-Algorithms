package leetcode

func IsHappy(n int) bool {
	values := make(map[int]bool)
	for n != 1 && !values[n] {
		values[n] = true
		x := 0
		for ; n > 0; n /= 10 {
			x += (n % 10) * (n % 10)
		}
		n = x
	}
	return n == 1
}
