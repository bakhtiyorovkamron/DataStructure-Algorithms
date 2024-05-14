package leetcode

func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	counter := 0
	rules := make(map[string]int)
	rules["color"] = 1
	rules["type"] = 0
	rules["name"] = 2
	for _, v := range items {
		if v[rules[ruleKey]] == ruleValue {
			counter++
		}
	}
	return counter
}
