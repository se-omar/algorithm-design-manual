package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	tree := &TreeNode{}

	if root1 == nil && root2 != nil {
		tree.Val = root2.Val
		tree.Left = mergeTrees(nil, root2.Left)
		tree.Right = mergeTrees(nil, root2.Right)
	}

	if root1 != nil && root2 == nil {
		tree.Val = root1.Val
		tree.Left = mergeTrees(root1.Left, nil)
		tree.Right = mergeTrees(root1.Right, nil)
	} else if root1 != nil && root2 != nil {
		tree.Val = root1.Val + root2.Val
		tree.Left = mergeTrees(root1.Left, root2.Left)
		tree.Right = mergeTrees(root1.Right, root2.Right)
	}

	return tree
}
