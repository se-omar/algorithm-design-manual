package main

func middleNode(head *ListNode) *ListNode {
	// 2 pointer approach
	slow := head
	fast := head

	for fast != nil && fast.Next != nil  {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}


// s    s    s
// O -> O -> O -> O -> O
// f         f         f
