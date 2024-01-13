package chapter3

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIter(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
