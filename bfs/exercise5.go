package bfs

import "fmt"

type Tree2 struct {
	Value int
	Left  *Tree2
	Right *Tree2
}

func FindTheMinInEveryNode(root *Tree2) {
	if root == nil {
		return
	}
	nodeMins := []int{}
	queue := []*Tree2{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		min := queue[0].Value
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Value < min {
				min = node.Value
			}
		}
		nodeMins = append(nodeMins, min)
	}
	fmt.Println("node mins :", nodeMins)
}
