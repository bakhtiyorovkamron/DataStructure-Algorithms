package main

import (
	"data-structure-algorithms/dfs"
)

func main() {
	root := &dfs.TreeNode{
		Val: 1,
	}
	root.Left = &dfs.TreeNode{Val: 2}
	root.Right = &dfs.TreeNode{Val: 3}
	root.Left.Left = &dfs.TreeNode{Val: 4}
	root.Left.Right = &dfs.TreeNode{Val: 5}
	root.Right.Left = &dfs.TreeNode{Val: 6}
	root.Right.Right = &dfs.TreeNode{Val: 7}
	dfs.SumTree(root)
}
