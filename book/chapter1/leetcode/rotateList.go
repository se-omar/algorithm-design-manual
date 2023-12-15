// Definition for singly-linked list.
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	tail := head
	curr := head
	length := 1
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	k = k % length
	if k == 0 {
		return head
	}

	shift := length - k - 1
	for i := 0; i < shift; i++ {
		curr = curr.Next
	}
	newHead := curr.Next
	curr.Next = nil
	tail.Next = head

	return newHead
}
