package main

var balanced bool
func isBalanced(root *TreeNode) bool {
	balanced = true
	helper(root)
	return balanced
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := helper(root.Left), helper(root.Right)
	if diff(lh, rh) > 1 {
		balanced = false
	}

	return  1 + max(lh, rh)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
