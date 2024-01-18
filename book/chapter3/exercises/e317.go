package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := treeHeight(root.Left)
	rh := treeHeight(root.Right)

	if diff(lh, rh) > 1 {
		return false
	}

	lb := isBalanced(root.Left)
	rb := isBalanced(root.Right)

	if lb == false || rb == false {
		return false
	}

	return true
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(treeHeight(root.Left), treeHeight(root.Right))
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
