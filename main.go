package main

import (
	"data-structure-algorithms/leetcode"
)

func main() {
	list := leetcode.ListNode{Val: 1}
	list.AddTail(2)
	list.AddTail(3)
	list.AddTail(4)
	list.AddTail(5)

	l := leetcode.MiddleNode(&list)
	l.Iterate()
	// n.Iterate()
}
