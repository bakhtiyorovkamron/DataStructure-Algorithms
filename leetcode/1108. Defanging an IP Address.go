package leetcode

func DefangIPaddr(address string) string {
    defanged := ""
	for _, c := range address {
		if c == '.' {
			defanged += "[.]"
		} else {
			defanged += string(c)
		}
	}
	return defanged
}