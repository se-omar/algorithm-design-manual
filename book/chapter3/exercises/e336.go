package main

func middleNode(head *ListNode) *ListNode {
	// traverse twice approach
	midNode := head
	length := listLength(head)
	mid := length / 2

	for i := 0; i < mid; i++ {
		midNode = midNode.Next
	}
	return midNode
}

func listLength(node *ListNode) int {
	if node == nil {
		return 0
	}

	return 1 + listLength(node.Next)
}
