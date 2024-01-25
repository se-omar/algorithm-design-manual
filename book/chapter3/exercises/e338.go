package main

func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftTail := dfs(root.Left)
	rightTail := dfs(root.Right)

	if leftTail != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightTail != nil {
		return rightTail
	}

	if leftTail != nil {
		return leftTail
	}

	return root
}
