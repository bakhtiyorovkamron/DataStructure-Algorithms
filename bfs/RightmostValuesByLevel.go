package bfs

import "fmt"

func RightmostValuesByLevel(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	j := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		mostRight := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			mostRight = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		j++
		if j%2 == 0 {
			fmt.Println(mostRight)

		}
	}
}
