package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	sorted := []int{}
	dfs(root, &sorted)

	for i := 0; i < len(sorted); i++ {
		if i > 0 && sorted[i] <= sorted[i-1] {
			return false
		}
	}

	return true
}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs(root.Right, arr)
}
