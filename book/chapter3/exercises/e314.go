package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	tree := &TreeNode{}
	sum := 0
	var leftNode1, leftNode2, rightNode1, rightNode2 *TreeNode = nil, nil, nil, nil

	if root1 != nil {
		sum += root1.Val
		leftNode1, rightNode1 = root1.Left, root1.Right
	}

	if root2 != nil {
		sum += root2.Val
		leftNode2, rightNode2 = root2.Left, root2.Right
	}
	tree.Val = sum
	tree.Left = mergeTrees(leftNode1, leftNode2)
	tree.Right = mergeTrees(rightNode1, rightNode2)

	return tree
}
