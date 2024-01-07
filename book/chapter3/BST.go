package chapter3

type BST struct {
	parent *BST
	left   *BST
	right  *BST
	item   int
}

func searchBst(tree *BST, item int) *BST {
	if tree == nil {
		return &BST{}
	}

	if item > tree.item {
		return searchBst(tree.right, item)
	} else if item < tree.item {
		return searchBst(tree.left, item)
	} else {
		return tree
	}
}

// finds the minimum element in a BST
func minBst(tree *BST) *BST {
	if tree == nil {
		return &BST{}
	}

	if tree.left == nil {
		return tree
	}

	return minBst(tree.left)
}

// finds the maximum element in a BST
func maxBst(tree *BST) *BST {
	if tree == nil {
		return &BST{}
	}

	if tree.right == nil {
		return tree
	}

	return maxBst(tree.right)
}

func traverseBst(tree *BST) {
	if tree != nil {
		processItem(tree.item)
		traverseBst(tree.left)
		traverseBst(tree.right)
	}
}

func processItem(item int) {
	// do something with each value from the tree traversal
}

func insertBst(tree *BST, item int, parent *BST) {
	if tree == nil {
		//insert item
		newTree := BST{parent: parent, left: nil, right: nil, item: item}
		*tree = newTree
		return
	}
	if item > tree.item {
		insertBst(tree.right, item, parent)
	} else if item < tree.item {
		insertBst(tree.left, item, parent)
	} else {
		return
	}

}
