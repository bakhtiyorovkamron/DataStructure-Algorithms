package leetcode

func NearestValidPoint(x int, y int, points [][]int) int {
	// index := 0
	mapValidPoints := make(map[int]int)
	min := 0
	for i := 0; i < len(points); i++ {
		if points[i][0] == x || points[i][1] == y {
			if _, ok := mapValidPoints[(Abs(x-points[i][0]) + Abs(y-points[i][1]))]; !ok {
				min = (Abs(x-points[i][0]) + Abs(y-points[i][1]))
				mapValidPoints[(Abs(x-points[i][0]) + Abs(y-points[i][1]))] = i
			}

		}
	}
	if len(mapValidPoints) == 0 {
		return -1
	}
	for k, _ := range mapValidPoints {
		if min > k {
			min = k
		}
	}
	return mapValidPoints[min]
}

