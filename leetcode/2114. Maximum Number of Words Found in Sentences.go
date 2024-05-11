package leetcode

func MostWordsFound(sentences []string) int {
    counter := 0 

	for _, sentence := range sentences {
		words := 1
		for _, c := range sentence {
			if c == ' ' {
				words++
			}
		}
		if words > counter {
			counter = words
		}
	}

	return counter
}