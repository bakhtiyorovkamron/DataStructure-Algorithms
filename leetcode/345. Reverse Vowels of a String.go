package leetcode

import (
	"strings"
)

func ReverseVowels(s string) string {

	starts, ends := []int{}, []int{}
	sI, eI := 0, 0
	target := []byte(s)
	k := len(target) - 1
	for i := 0; i < len(target); i++ {
		if isVowel(string(target[i])) {
			if len(ends) > 0 {
				target[i] = target[eI]
			}
			starts = append(starts, i)
			sI++
		}
		if isVowel(string(target[k])) {
			if len(starts) > 0 {
				target[k] = target[sI]
			}
			ends = append(ends, k)
			eI++
		}

	}

	// stringByte := []byte(s)
	// array := []uint8{}
	// for i := 0; i < len(stringByte); i++ {
	// 	if isVowel(string(stringByte[i])) {
	// 		array = append(array, stringByte[i])
	// 	}
	// }

	// k := 0
	// array = reverseArray(array)
	// for i := 0; i < len(stringByte); i++ {
	// 	if isVowel(string(stringByte[i])) {
	// 		stringByte[i] = array[k]
	// 		k++
	// 	}
	// }
	return "test result"
}
func isVowel(s string) bool {
	s = strings.ToLower(s)
	if s == "a" || s == "e" || s == "i" || s == "u" || s == "o" {
		return true
	}
	return false
}
func reverseArray(array []uint8) []uint8 {
	newArray := []uint8{}
	for i := len(array) - 1; i >= 0; i-- {
		newArray = append(newArray, array[i])
	}
	return newArray
}
