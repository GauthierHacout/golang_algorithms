package algorithms

//Definition for a binary tree Node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

// Create a BST from a slice
func BSTfromSlice(keys []int) *TreeNode{
	var root *TreeNode
	for _, key := range keys {
		root = insert(root, key)
	}
	return root
}

func newNode(x int) *TreeNode {
	return &TreeNode{Val:x, Left:nil, Right:nil}
}

func insert(root *TreeNode, x int) *TreeNode {
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
