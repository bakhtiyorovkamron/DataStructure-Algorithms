package dfs

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Println(root.Val, SumTree(root.Left), SumTree(root.Right))
	return root.Val + SumTree(root.Left) + SumTree(root.Right)
}
