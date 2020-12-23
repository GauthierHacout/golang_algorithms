package algorithms

import "errors"

// Return the K th smallest element that is inside a Binary Search Tree
func KthSmallest(root *TreeNode, k int) (int, error) {
	smallest, _ := searchKthSmallest(root, k, 0)
	if smallest != nil {
		return smallest.Val, nil
	}
	return 0, errors.New("Not enough elements inside BST")
}

func searchKthSmallest(root *TreeNode, k int, count int) (*TreeNode, int){
	// Using Inorder Traversal
	// Base case
	if (root == nil){
		return nil, count
	}

	// Search in left subtree
	left, count := searchKthSmallest(root.Left, k, count)

	// If k'th smallest is found in left subtree return it
	if (left != nil) {
		return left, count
	}

	// If current element is k'th smallest, return it
	count++;
	if (count==k) {
		return root, count
	}

	//Else search in the right subtree
	return searchKthSmallest(root.Right, k, count)
}