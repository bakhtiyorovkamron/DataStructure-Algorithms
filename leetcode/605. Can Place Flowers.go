package leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	noll := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			noll++
		}
	}

	return noll >= n*2
}
