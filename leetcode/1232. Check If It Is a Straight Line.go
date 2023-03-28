package leetcode

import "fmt"

func CheckStraightLine(coordinates [][]int) bool {
	a1 := coordinates[1][0] - coordinates[0][0]
	a2 := coordinates[1][1] - coordinates[0][1]
	for i := 0; i < len(coordinates); i++ {
		if i != len(coordinates)-1 {
			b1 := coordinates[i+1][0] - coordinates[i][0]
			b2 := coordinates[i+1][1] - coordinates[i][0]
			fmt.Println(b2, b1)
			if b1*a2 != b2*a1 {
				return false
			}
		}
	}
	return true
}

func partition1(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func quickSort1(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition1(arr, low, high)
		arr = quickSort1(arr, low, p-1)
		arr = quickSort1(arr, p+1, high)
	}
	return arr
}

func quickSortStart1(arr []int) []int {
	return quickSort1(arr, 0, len(arr)-1)
}
