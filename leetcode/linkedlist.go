package leetcode

import "fmt"

type List struct {
	count int
	head  *Node
}
type Node struct {
	value int
	next  *Node
}

func (l *List) Size() int {
	return l.count
}
func (l *List) IsEmpty() bool {
	return l.count == 0
}
func (l *List) AddHead(value int) {
	l.head = &Node{value, l.head}
	l.count++
}
func (l *List) Iterate() {
	for l.head != nil {
		fmt.Println(l.head.value)
		l.head = l.head.next
	}
}
func (l *List) AddTail(value int) {
	curr := l.head
	newNode := &Node{value, nil}
	if curr == nil {
		l.head = newNode
		return
	}
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}
