package leetcode

func ReplaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			firsRightElement := arr[i+1]
			for j := i + 1; j < len(arr); j++ {
				if arr[j] > firsRightElement {
					firsRightElement = arr[j]
				}
			}
			arr[i] = firsRightElement
		} else {
			arr[i] = -1
		}
	}
	return arr
}
