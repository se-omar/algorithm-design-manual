package main

func middleNode(head *ListNode) *ListNode {
	// 2 pointer approach
	slow := head
	fast := head

	for {
		if fast == nil || fast.Next == nil {
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
