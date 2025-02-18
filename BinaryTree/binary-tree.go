package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func insertNode(root *TreeNode, newValue int) *TreeNode {

	// Create a new node if the tree is empty
	if root == nil {
		return &TreeNode{value: newValue}
	}

	if newValue < root.value {
		root.left = insertNode(root.left, newValue)
	} else {
		root.right = insertNode(root.right, newValue)
	}
	return root
}

func main() {
	var treeRoot *TreeNode

	treeRoot = insertNode(treeRoot, 5)
	treeRoot = insertNode(treeRoot, 3)
	treeRoot = insertNode(treeRoot, 7)
	treeRoot = insertNode(treeRoot, 2)
	treeRoot = insertNode(treeRoot, 4)
	treeRoot = insertNode(treeRoot, 6)
	treeRoot = insertNode(treeRoot, 8)
}
