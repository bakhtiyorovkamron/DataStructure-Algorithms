package leetcode

func BackspaceCompare(s string, t string) bool {
	ssharp, tsharp := "", ""
	ssharpCount, tsharpCount := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != "#" && ssharpCount == 0 {
			ssharp += string(s[i])
		}
		if string(s[i]) != "#" && ssharpCount != 0 {
			ssharpCount--
		}
		if string(s[i]) == "#" {
			ssharpCount++
		}
	}
	for i := len(t) - 1; i >= 0; i-- {
		if string(t[i]) != "#" && tsharpCount == 0 {
			tsharp += string(t[i])
		}
		if string(t[i]) != "#" && tsharpCount != 0 {
			tsharpCount--
		}
		if string(t[i]) == "#" {
			tsharpCount++
		}
	}
	return ssharp == tsharp
}
