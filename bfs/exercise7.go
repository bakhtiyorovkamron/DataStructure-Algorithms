package bfs

import "fmt"

type Tree4 struct {
	Val   int
	Left  *Tree4
	Right *Tree4
}

func FindSumOfLeafNodes(root *Tree4) {
	if root == nil {
		return
	}

	queue := []*Tree4{root}
	sum := 0
	for len(queue) > 0 {

		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			left := node.Left
			if left != nil {
				queue = append(queue, left)
			}
			right := node.Right
			if right != nil {
				queue = append(queue, right)
			}
			if right == nil && left == nil {
				sum += node.Val
			}

		}
	}
	fmt.Println(sum)
}
