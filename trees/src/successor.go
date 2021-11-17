package src

import (
	"bst/adt"
	"fmt"
)

func Successor(root *adt.NodeWithParentPointer, element int ) *adt.NodeWithParentPointer {
	if root == nil {
		return nil
	}

	if root.Val > element {
		return Successor(root.Left,element)
	}

	if root.Val == element {
	current := root

	if current.Right != nil {
		//go left most
		leftMost := getLeftMostElement(current.Right)

		return leftMost
	}

	if current.Right == nil {
		//left child
		current = current.Parent
		//right child
		fmt.Println("Current.Right : ",current.Right.Val)
		if current.Right.Val == element {
			sucessor := traverseRight(current)
			fmt.Println("Sucessor :",sucessor)
			return sucessor
		} 
		return current
	}

}
	return Successor(root.Right,element)
}

func getLeftMostElement(root *adt.NodeWithParentPointer) *adt.NodeWithParentPointer{

	if root.Left == nil {
		return root
	}

	fmt.Println("Root.Left : ",root.Left)
	return getLeftMostElement(root.Left)
}

func traverseRight(root *adt.NodeWithParentPointer) *adt.NodeWithParentPointer {
	
	current := root.Parent
	if current != nil {
		traverseRight(root.Parent)
	} 

	return current
}