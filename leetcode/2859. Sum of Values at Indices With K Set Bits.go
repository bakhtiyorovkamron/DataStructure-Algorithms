package leetcode

import (
	"fmt"
	"strings"
)

func decimalToBinary(num int) string {
	var binary []int
	bin := ""

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}
	if len(binary) == 0 {
		return fmt.Sprintf("%d\n", 0)
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			bin += fmt.Sprintf("%d", binary[i])
		}

	}
	return bin
}

func SumIndicesWithKSetBits(nums []int, k int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if strings.Count(decimalToBinary(i), "1") == k {
			counter += nums[i]
		}
	}
	return counter
}
