package leetcode

func FindSubstring(s string, k int32) string {
	// Write your code here

	vowels := []string{}

	count := 0
	for _, v := range s {
		if string(v) == "a" || string(v) == "e" || string(v) == "i" || string(v) == "o" || string(v) == "u" {
			count++
		}
	}
	if count == 0 {
		return "Not found!"

	}
	result := make(map[string]int)

	for i := 0; i < len(s); i++ {
		str := ""
		for j := i; j < i+int(k); j++ {
			if j == len(s) {
				break
			}
			str += string(s[j])
		}
		if len(str) == int(k) {
			vowels = append(vowels, str)
			count := 0
			for s := 0; s < len(str); s++ {
				if string(str[s]) == "a" || string(str[s]) == "e" || string(str[s]) == "i" || string(str[s]) == "o" || string(str[s]) == "u" {
					count++
				}
			}
			result[str] = count
		}
	}
	max := 0
	resultstring := ""
	for i, v := range result {
		if max < v {
			max = v
			resultstring = i
		}
	}
	if resultstring != "" {
		return resultstring
	}

	return "Not found!"
}
func ReturninStrings(str string, k int32) []string {
	arr := []string{}
	for i := 0; i < len(str); i++ {
		s := ""
		for j := i; j < i+int(k); j++ {
			if j == len(str) {
				return arr
			}
			s += string(str[j])
		}
		if len(s) == int(k) {
			arr = append(arr, s)
		}
	}
	return arr
}
