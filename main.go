package main

import (
	"data-structure-algorithms/leetcode"
	"fmt"
)

func main() {
	// tree := leetcode.NodeN{
	// 	Val: 1,
	// 	Children: []*leetcode.NodeN{
	// 		{
	// 			Val: 3,
	// 			Children: []*leetcode.NodeN{
	// 				{
	// 					Val:      5,
	// 					Children: []*leetcode.NodeN{},
	// 				},
	// 				{
	// 					Val:      6,
	// 					Children: []*leetcode.NodeN{},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Val:      2,
	// 			Children: []*leetcode.NodeN{},
	// 		},
	// 		{
	// 			Val:      4,
	// 			Children: []*leetcode.NodeN{},
	// 		},
	// 	},
	// }
	// (leetcode.Preorder(&tree))
	//leetcode.NextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7})
	fmt.Println(leetcode.Interpret("G()()()()(al)"))
}

// [1,3,5,6,2,4]
