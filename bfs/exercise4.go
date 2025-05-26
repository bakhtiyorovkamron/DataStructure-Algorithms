package bfs

import "fmt"

func Exercise4(root *Tree) {
	if root == nil {
		return
	}
	var result []int
	queue := []*Tree{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		result = append(result, levelSize)
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
	}
	fmt.Println("LevelSize :", result)
}
