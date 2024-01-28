package main

type LinkedList struct {
	item any
	next *LinkedList
}

func hasCycle(head *LinkedList) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}

	return false
}
