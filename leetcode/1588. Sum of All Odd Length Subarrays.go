package leetcode

func SumOddLengthSubarrays(arr []int) int {
	sum := 0
	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j < j+i; j++ {
			if j+i <= len(arr) {
				for _, v := range arr[j : j+i] {
					sum += v
				}
			}
		}
	}
	
	return sum
}

// 1, 4, 2, 5, 3
