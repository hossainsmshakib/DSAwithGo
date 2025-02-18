package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Store the diameter path
var diameterPath []int

// Insert a new node into the Binary Search Tree
func insertNode(root *TreeNode, newValue int) *TreeNode {
	// Create a new node if the tree is empty
	if root == nil {
		return &TreeNode{value: newValue}
	}

	// Insert in the left or right subtree based on value
	if newValue < root.value {
		root.left = insertNode(root.left, newValue)
	} else {
		root.right = insertNode(root.right, newValue)
	}
	return root
}

// Search for a node in the Binary Search Tree
func searchNode(root *TreeNode, targetValue int) bool {
	if root == nil {
		return false
	}

	// If found, return true
	if root.value == targetValue {
		return true
	}

	// Recur to left or right subtree
	if targetValue < root.value {
		return searchNode(root.left, targetValue)
	} else {
		return searchNode(root.right, targetValue)
	}
}

// Delete a node from the Binary Search Tree
func deleteNode(root *TreeNode, targetValue int) *TreeNode {
	// If tree is empty
	if root == nil {
		return nil
	}

	// Find the node to delete
	if targetValue < root.value {
		root.left = deleteNode(root.left, targetValue)
	} else if targetValue > root.value {
		root.right = deleteNode(root.right, targetValue)
	} else {
		// Found the node to delete

		// Case 1: Node has no child
		if root.left == nil && root.right == nil {
			return nil // Remove the node
		}

		// Case 2: Node has only one child
		if root.left == nil {
			return root.right // Replace with right child
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

// Find the node with the smallest value in a subtree
func findMinValue(root *TreeNode) *TreeNode {
	if root.left == nil {
		return root // The leftmost node is the smallest
	}
	return findMinValue(root.left) // Keep going left
}

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

// Compute the height of the Binary Search Tree
func calculateTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0 // If tree is empty, height is 0
	}
	leftSubtreeHeight := calculateTreeHeight(root.left)
	rightSubtreeHeight := calculateTreeHeight(root.right)
	return 1 + max(leftSubtreeHeight, rightSubtreeHeight) // Max depth of left/right subtrees
}

// Reverse (Mirror) a Binary Tree
func reverseTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right subtrees
	root.left, root.right = root.right, root.left

	// Recursively reverse left and right subtrees
	reverseTree(root.left)
	reverseTree(root.right)

	return root
}

// Compute the Diameter of the Binary Tree
func findDiameter(root *TreeNode) (int, []int) {
	if root == nil {
		return 0, []int{} // Height is 0, path is empty
	}

	// Get height & path of left and right subtrees
	leftHeight, leftPath := findDiameter(root.left)
	rightHeight, rightPath := findDiameter(root.right)

	// Path that goes through the root
	currentPath := append(leftPath, root.value)
	currentPath = append(currentPath, rightPath...)

	// Determine the longest path
	if leftHeight > rightHeight {
		return leftHeight + 1, append(leftPath, root.value)
	} else {
		return rightHeight + 1, append(rightPath, root.value)
	}
}

// Compute the Actual Diameter and Track the Path
func getDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight, leftPath := findDiameter(root.left)
	rightHeight, rightPath := findDiameter(root.right)

	// Path that goes through the root
	currentPath := append(leftPath, root.value)
	currentPath = append(currentPath, rightPath...)

	// Update global diameter path if the new one is longer
	if len(currentPath) > len(diameterPath) {
		diameterPath = currentPath
	}

	// Return the maximum diameter found
	return max(leftHeight+rightHeight, max(getDiameter(root.left), getDiameter(root.right)))
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

	fmt.Println("Preorder Traversal:")
	preorderTraversal(treeRoot)
	fmt.Println("\nInorder Traversal:")
	inorderTraversal(treeRoot)
	fmt.Println("\nPostorder Traversal:")
	postorderTraversal(treeRoot)

	fmt.Println("\nHeight of Tree:", calculateTreeHeight(treeRoot))

	// Search
	fmt.Println("\nSearch for 4:", searchNode(treeRoot, 4))
	fmt.Println("Search for 10:", searchNode(treeRoot, 10))

	// Delete node
	treeRoot = deleteNode(treeRoot, 7)
	fmt.Println("\nInorder Traversal After Deletion:")
	inorderTraversal(treeRoot)
	fmt.Println()

	// Reverse (Mirror) the Binary Tree
	treeRoot = reverseTree(treeRoot)
	fmt.Println("\nInorder Traversal After Reversing Tree:")
	inorderTraversal(treeRoot) // The tree should now be mirrored
	fmt.Println()

	// Compute Diameter and Display Path
	fmt.Println("\nDiameter of Tree:", getDiameter(treeRoot))
	fmt.Println("Diameter Path:", diameterPath)
}
