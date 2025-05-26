package bfs

import "fmt"

func BFSFindShallowestLeafDepth(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	depth := 1
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				fmt.Println(depth)
				return
			}
		}
		depth++
	}

}
