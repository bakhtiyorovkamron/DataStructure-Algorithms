package leetcode

import (
	"math"
)

func ScoreOfString(s string) int {
	sum := 0
	for i, v := range s {
		if i != len(s)-1 {
			sum += int(math.Abs(float64(int(v) - int(s[i+1]))))
		}
	}
	return sum
}
