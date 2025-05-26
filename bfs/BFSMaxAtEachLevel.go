package bfs

import "fmt"

func BFSMaxAtEachLevel(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	maxes := []int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		max := queue[0].Val
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		maxes = append(maxes, max)
	}
	fmt.Println(maxes)
}
