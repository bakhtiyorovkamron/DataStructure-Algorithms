package leetcode

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 || len(mat)*len(mat[0]) != r*c || len(mat) == r && len(mat[0]) == c {
		return mat
	}

	res := make([][]int, r)
	count, col := 0, len(mat[0])
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = mat[count/col][count%col]
			count++
		}
	}

	return res
}
