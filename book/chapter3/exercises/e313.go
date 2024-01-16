package main

import "math"

var first, mid, last, prev *TreeNode

func recoverTree(root *TreeNode) {
	first, mid, last, prev = nil, nil, nil, &TreeNode{Val: math.MinInt}
	inorder(root)

	if first != nil && last != nil {
		first.Val, last.Val = last.Val, first.Val
	} else if first != nil && mid != nil {
		first.Val, mid.Val = mid.Val, first.Val
	}

}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)

	if prev != nil && root.Val < prev.Val {
		if first == nil {
			first = prev
			mid = root
		} else {
			last = root
		}
	}
	prev = root

	inorder(root.Right)
}
