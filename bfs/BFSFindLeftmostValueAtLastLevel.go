package bfs

import "fmt"

func BFSFindLeftmostValueAtLastLevel(root *TreeNode) {
	// if root == nil {
	// 	return
	// }
	// queue := []*TreeNode{root}
	// leftmost := 0
	// for len(queue) > 0 {
	// 	node := queue[0]
	// 	queue = queue[1:]
	// 	if node.Left != nil {
	// 		leftmost = node.Left.Val
	// 		queue = append(queue, node.Left)
	// 	}
	// 	if node.Right != nil {
	// 		queue = append(queue, node.Right)
	// 	}
	// }
	// fmt.Println(leftmost)
	queue := []*TreeNode{root}
	var theMostLeftNode int

	for len(queue) > 0 {
		lvlSize := len(queue)
		theMostLeftNode = queue[0].Val
		for i := 0; i < lvlSize; i++ {
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
	fmt.Println(theMostLeftNode)
}
