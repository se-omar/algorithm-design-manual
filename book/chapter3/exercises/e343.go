package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	ht := make(map[*ListNode]int)

	for head != nil {
		if _, ok := ht[head]; ok {
			return true
		}
		ht[head] = head.Val
		head = head.Next
	}
	return false
}
