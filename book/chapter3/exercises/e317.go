package main

type treeInfo struct {
	balanced bool
	h        int
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	info := helper(root, 0)
	return info.balanced
}

func helper(root *TreeNode, h int) treeInfo {
	if root == nil {
		return treeInfo{balanced: true, h: 0}
	}

	left, right := helper(root.Left, h+1), helper(root.Right, h+1)
	balanced := left.balanced && right.balanced && diff(left.h, right.h) <= 1
	return treeInfo{balanced: balanced, h: 1 + max(left.h, right.h)}
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
