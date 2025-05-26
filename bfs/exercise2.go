package bfs

import "fmt"

func FindDepth(root *Tree) {
	queue := []*Tree{root}
	depth := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		fmt.Println("levelsize :", levelSize)
		for i := 0; i < levelSize; i++ {
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
