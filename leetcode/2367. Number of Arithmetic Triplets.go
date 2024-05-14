package leetcode

import "fmt"

func ArithmeticTriplets(nums []int, diff int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				fmt.Println(i, j)
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					counter++
				}
			}
		}
	}
	fmt.Println("result :", counter)
	return counter
}
