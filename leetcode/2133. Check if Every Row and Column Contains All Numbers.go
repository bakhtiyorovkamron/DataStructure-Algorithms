package leetcode

func CheckValid(matrix [][]int) bool {
	checkMap := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			checkMap[matrix[i][j]] += 1
		}
	}
	for _, v := range checkMap {
		if v != len(matrix) {
			return false
		}
	}
	return true
}
