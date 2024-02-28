package leetcode

func FindMaxAverage(nums []int, k int) float64 {
	// 1, 12, -5, -6, 50, 3
	// for i := 0; i <= len(nums)-k; i++ {
	// 	fmt.Println(nums[i : i+k])
	// }
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	m := max
	l := len(nums)
	i := k
	for i < l {
		max = (max - nums[i-k]) + nums[i]
		if max > m {
			m = max
		}
		i++
	}
	return float64(m) / float64(k)
	// return 0
}
