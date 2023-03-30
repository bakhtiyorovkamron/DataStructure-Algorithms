package leetcode

func SumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	ans := 0
	for i := range arr {
		s := 0
		for j := i; j < n; j++ {
			s += arr[j]
			if (j-i+1)%2 == 1 {
				ans += s
			}
		}
	}
	return ans
}
// 1, 4, 2, 5, 3
