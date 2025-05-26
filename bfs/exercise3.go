package bfs

import "fmt"

func FindSumAllLevel(root *Tree) {

	if root == nil {
		return
	}

	queue := []*Tree{root}
	sum := []int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		innerSum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			innerSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		sum = append(sum, innerSum)

	}
	fmt.Println(sum)
}
