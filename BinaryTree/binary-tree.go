package main

import "fmt"

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

func searchNode(root *TreeNode, targetValue int) bool {
	if root == nil {
		return false
	}

	if root.value == targetValue {
		return true
	}

	if targetValue < root.value {
		return searchNode(root.left, targetValue)
	} else {
		return searchNode(root.right, targetValue)
	}
}

func deleteNode(root *TreeNode, targetValue int) *TreeNode {

	// tree is empty
	if root == nil {
		return nil
	}
	// find the node to del
	if targetValue < root.value {
		root.left = deleteNode(root.left, targetValue)
	} else if targetValue > root.value {
		root.right = deleteNode(root.right, targetValue)
	} else {
		// found the node now del it

		// case 1: Node has no child
		if root.left == nil && root.right == nil {
			return nil // remove the node
		}

		// Case 2: Node has only one child
		if root.left == nil {
			return root.right // replace with right child
		}
		if root.right == nil {
			return root.left
		}
		// Case 3: Node has two children
		// Find the smallest value in the right subtree (inorder successor)
		smallestNode := findMinValue(root.right)
		root.value = smallestNode.value
		root.right = deleteNode(root.right, smallestNode.value) // Delete the inorder successor
	}
	return root
}

/*
Step-by-Step Process
Find the inorder successor

The inorder successor is the smallest value in the right subtree of the node to be deleted.
It is the leftmost node in the right subtree.
Replace the value of the node to be deleted

Instead of deleting the node directly, copy the value of the inorder successor into the current node.
Delete the inorder successor node

Since the inorder successor node is the smallest node in the right subtree, it cannot have a left child.
We delete it from the tree.
*/

// Find the node with the smallest value (Helper function for deletion)
func findMinValue(root *TreeNode) *TreeNode {
	if root.left == nil {
		return root // The leftmost node is the smallest
	}
	return findMinValue(root.left) // Keep going left
}

/*
We go left because the smallest value is always in the leftmost node of a Binary Search Tree.
✅ Left child always has smaller values
✅ Going right would lead to larger values
*/

// Preorder Traversal (Root -> Left -> Right)
func preorderTraversal(root *TreeNode) {
	if root != nil {
		fmt.Print(root.value, " ")
		preorderTraversal(root.left)
		preorderTraversal(root.right)
	}
}

// Inorder Traversal (Left -> Root -> Right)
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.left)
		fmt.Print(root.value, " ")
		inorderTraversal(root.right)
	}
}

// Postorder Traversal (Left -> Right -> Root)
func postorderTraversal(root *TreeNode) {
	if root != nil {
		postorderTraversal(root.left)
		postorderTraversal(root.right)
		fmt.Print(root.value, " ")
	}
}

// Compute the height of the Binary Search Tree (Recursive)
func calculateTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0 // If tree is empty, height is 0
	}
	leftSubtreeHeight := calculateTreeHeight(root.left)
	rightSubtreeHeight := calculateTreeHeight(root.right)
	return 1 + max(leftSubtreeHeight, rightSubtreeHeight) // Max depth of left/right subtrees
}

// Helper function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
