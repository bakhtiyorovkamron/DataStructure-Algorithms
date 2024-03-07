package leetcode

func MiddleNode2(head *ListNode) *ListNode {
	var (
		sz, counter int
		// newList     *ListNode
		temp = head
	)

	for temp.Next != nil {
		sz++
		temp = temp.Next
	}
	if sz == 0 {
		return head
	}
	if sz == 1 {
		head = head.Next
		return head
	}
	sz = (sz - 1) / 2
	for head.Next != nil {
		if counter > sz {
			return head
		}
		counter++

		head = head.Next

	}
	return &ListNode{}
}
