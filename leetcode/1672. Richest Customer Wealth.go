package leetcode

func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for i := 0; i < len(accounts[0]); i++ {
		maxWealth += accounts[0][i]
	}
	for i := 0; i < len(accounts); i++ {
		wealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			wealth += accounts[i][j]
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
