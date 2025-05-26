package bfs

import "fmt"

func PrintEvenLevelSums(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		evenSum := 0
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			evenSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if evenSum%2 == 0 {
			fmt.Println(evenSum)
		}
	}
}
