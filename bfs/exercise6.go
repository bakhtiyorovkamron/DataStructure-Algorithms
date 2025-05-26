package bfs

import "fmt"

type Tree3 struct {
	Value int
	Chap  *Tree3
	Ong   *Tree3
}

func FindArifmetic(root *Tree3) {
	if root == nil {
		return
	}
	queue := []*Tree3{root}
	result := []float32{}
	for len(queue) > 0 {
		currentLength := len(queue)
		var sum float32 = 0
		for i := 0; i < currentLength; i++ {
			node := queue[0]
			sum += float32(node.Value)
			queue = queue[1:]
			if node.Chap != nil {
				queue = append(queue, node.Chap)
			}
			if node.Ong != nil {
				queue = append(queue, node.Ong)
			}
		}
		result = append(result, sum/float32(currentLength))
	}
	fmt.Println(result)
}
