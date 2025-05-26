package bfs

import "fmt"

func BFSPrintByLevels(root *TreeNode) {
	if root == nil {
		return
	}
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		level := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	fmt.Println(result)
}
