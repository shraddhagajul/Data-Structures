package src

import (
	"bst/adt"
)

func BstSearching(root *adt.Node, element int) *adt.Node {

	if root == nil{
		return nil
	}

	if element == root.Val{
		return root
	}

	if element < root.Val {
		return BstSearching(root.Left,element)
	}

	if element > root.Val{
		return BstSearching(root.Right,element)
	}

	return nil

}
