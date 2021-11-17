package src

import "bst/adt"


func IsBst(root *adt.Node) bool {

	if root == nil {
		return true
	}
	
	if root.Left != nil && root.Left.Val > root.Val{
		return false
	}

	if root.Right != nil && root.Right.Val < root.Val {
		return false
	}

	return IsBst(root.Left) &&  IsBst(root.Right)
	
}