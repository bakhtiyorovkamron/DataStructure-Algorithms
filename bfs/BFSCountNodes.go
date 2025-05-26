package bfs

import "fmt"

func BFSCountNodes(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	counter := 0
	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]
		counter++

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	fmt.Println("counter :", counter)
}
