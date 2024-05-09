package leetcode

func FinalValueAfterOperations(operations []string) int {
	counter := 0
    for _, op := range operations {
		if op == "++X" || op == "X++" {
			counter++
		} else {
			counter--
		}
	}
	return counter
}