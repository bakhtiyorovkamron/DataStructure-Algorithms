package leetcode

func DiagonalSum(mat [][]int) int {
	i := 0
	middle := 0
	counter := 0
	for _, v := range mat {
		counter += v[i]
		i++
	}
	if i%2 != 0 {
		middle = mat[i/2][i/2]
	}
	for _, v := range mat {
		i--
		counter += v[i]
	}
	return counter - middle
}
