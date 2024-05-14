package leetcode

func LargestAltitude(gain []int) int {
	altitude := make([]int, len(gain)+1)

	for i := 1; i <= len(gain); i++ {
		altitude[i] = altitude[i-1] + gain[i-1]
	}
	theBiggest := altitude[0]
	for _, v := range altitude {
		if v > theBiggest {
			theBiggest = v
		}
	}

	return theBiggest
}
