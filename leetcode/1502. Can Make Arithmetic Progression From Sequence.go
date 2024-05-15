package leetcode

func CanMakeArithmeticProgression(arr []int) bool {
	arr = quickSortStart(arr)
	progression := arr[1] - arr[0]
	arrNull := arr[0]
	for i := 0; i < len(arr); i++ {
		if arrNull != arr[i] {
			return false
		}
		arrNull += progression
	}
	return true
}

func partition(arr []int, low, high int) ([]int, int) {
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
// func quickSort(arr []int, low, high int) []int {
// 	if low < high {
// 		var p int
// 		arr, p = partition(arr, low, high)
// 		arr = quickSort(arr, low, p-1)
// 		arr = quickSort(arr, p+1, high)
// 	}
// 	return arr
// }

// func quickSortStart(arr []int) []int {
// 	return quickSort(arr, 0, len(arr)-1)
// }
