package bfs

import "fmt"

func BFSFindDeepestLeafDepth(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		levl := len(queue)
		for i := 0; i < levl; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	fmt.Println(depth)
}
