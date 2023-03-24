package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) AddTail(value int) {
	curr := l
	newNode := &ListNode{value, nil}
	if curr == nil {
		curr = newNode
		return
	}
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}
func (l *ListNode) AddHead(value int) {
	l.Next = &ListNode{value, l.Next}
}
func (list *ListNode) Print() {
	temp := list
	for temp != nil {
		fmt.Print(temp.Val, " ")
		temp = temp.Next
	}
	fmt.Println("")
}
func (list *List) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}
func DeleteDuplicates(head *ListNode) *ListNode {
	curr := head
	if curr == nil {
		return nil
	}
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
func MiddleNode(head *ListNode) *ListNode {
	curr := head
	curr2 := head
	if curr == nil {
		return nil
	}
	counter := 0
	for curr != nil {
		curr = curr.Next
		counter++
	}
	middle := counter / 2
	counter = 0
	arr := []int{}
	for curr2 != nil {
		if counter >= middle {
			arr = append(arr, curr2.Val)
		}
		curr2 = curr2.Next

		counter++
	}
	newNode := &ListNode{arr[len(arr)-1], nil}
	for i := len(arr) - 2; i >= 0; i-- {
		newNode = &ListNode{arr[i], newNode}
	}
	return newNode
}
func (list *List) Reverse() {
	curr := list.head
	var next, prev *Node
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}
func IsPalindromeNode(head *ListNode) bool {
	str := "opo"
	for i := 0 ; i < len(str) ; i++ {
		if string(str[i]) != string(str[len(str)-1-i]) {
			return false
		}
	}
	return true
}
