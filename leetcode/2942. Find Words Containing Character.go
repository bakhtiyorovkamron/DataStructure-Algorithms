package leetcode


// import "fmt"

func FindWordsContaining(words []string, x byte) []int {
	// fmt.Println(words)
	indices := []int{}
	for i,v := range words {
		isExist := false
		for j := 0 ;j < len(v);j++ {
			if v[j]==x {
				isExist = true
			}
		}
		if isExist {
			indices=append(indices,i)
		}
	}

	return indices
    
}