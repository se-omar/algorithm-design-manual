package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		if fast.Next == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
