package chapter3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	countL := 1 + maxDepth(root.Left)
	countR := 1 + maxDepth(root.Right)

	if countL > countR {
		return countL
	}

	return countR

}
