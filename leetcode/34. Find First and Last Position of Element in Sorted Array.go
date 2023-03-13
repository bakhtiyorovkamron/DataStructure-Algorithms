package leetcode

import "fmt"

func SearchRange(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	leftIndex, rightIndex := -1, -1
	fmt.Println(nums)
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			fmt.Println("mid", mid)
			leftIndex = mid
			j = mid - 1
		}
	}

	i, j = 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)/2

		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			rightIndex = mid
			i = mid + 1
		}
	}

	return []int{leftIndex, rightIndex}

}
