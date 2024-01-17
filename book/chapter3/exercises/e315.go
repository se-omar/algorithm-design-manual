package main

func convertBST(root *TreeNode) *TreeNode {
	var sortedArr []int
	inorder2(root, sortedArr)
	return createBST(sortedArr)
}

func inorder2(root *TreeNode, sortedArr []int) []int {
	if root == nil {
		return []int{}
	}

	inorder2(root.Left, sortedArr)
	sortedArr = append(sortedArr, root.Val)
	inorder2(root.Right, sortedArr)
	return sortedArr
}

func createBST(sortedArr []int) *TreeNode {
	var tree *TreeNode

	if len(sortedArr) <= 0 {
		return nil
	}

	mid := len(sortedArr) / 2
	tree.Val = sortedArr[mid]

	tree.Left = createBST(sortedArr[:mid])
	tree.Right = createBST(sortedArr[mid+1:])

	return tree
}
