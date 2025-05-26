package leetcode

func Construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	result := make([][]int, m)
	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i] = append(result[i], original[k])
			k++
		}
	}
	return result
}
