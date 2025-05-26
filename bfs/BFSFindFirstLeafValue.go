package bfs

import "fmt"

func BFSFindFirstLeafValue(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levlSize := len(queue)
		for i := 0; i < levlSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Right == nil && node.Left == nil {
				fmt.Println(node.Val)
				return
			}
		}
	}
}
