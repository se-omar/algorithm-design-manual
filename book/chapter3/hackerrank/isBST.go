package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBst(root *TreeNode) bool {
	if root == nil {
		return true
	}

	arr := []int{}
	bstToArr(root, &arr)

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func bstToArr(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	bstToArr(root.Left, arr)
	*arr = append(*arr, root.Val)
	bstToArr(root.Right, arr)
}
