package algorithms

//Definition for a binary tree node.
type TreeNode struct {
  Val byte
  Left *TreeNode
  Right *TreeNode
}

// Create a BST from a slice
func BSTfromSlice(keys []byte) *TreeNode{
	var root *TreeNode
	for _, key := range keys {
		root = insert(root, key)
	}
	return root
}

func newNode(x byte) *TreeNode {
	return &TreeNode{Val:x, Left:nil, Right:nil}
}

func insert(root *TreeNode, x byte) *TreeNode {
	if root==nil {
		return newNode(x)
	}

	if x < root.Val {
		root.Left = insert(root.Left, x)
	} else if x > root.Val {
		root.Right = insert(root.Right, x)
	}
	return root;
}
