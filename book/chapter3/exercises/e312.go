package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthRecur(root.Left), maxDepthRecur(root.Right))
}

type stack []stackItem

type stackItem struct {
	node  *TreeNode
	level int
}

func (s *stack) pop() stackItem {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *stack) push(el stackItem) {
	*s = append(*s, el)
}

func maxDepthIter(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLevel := 1
	stack := stack{stackItem{root, 1}}

	for len(stack) > 0 {
		newItem := stack.pop()
		if newItem.node == nil {
			continue
		}

		maxLevel = max(maxLevel, newItem.level)
		stack.push(stackItem{newItem.node.Left, newItem.level + 1})
		stack.push(stackItem{newItem.node.Right, newItem.level + 1})
	}

	return maxLevel
}
