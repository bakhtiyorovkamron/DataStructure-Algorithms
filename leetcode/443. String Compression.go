package leetcode

func Compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	ch := make(map[byte]int)
	for _, c := range chars {
		ch[c]++
	}

	return len(ch) * 2
}
