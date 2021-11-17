package src

import (
	"bst/adt"
)


func LcaInBinaryTree(root *adt.NodeWithParentPointer, first, second *adt.NodeWithParentPointer, parent *adt.NodeWithParentPointer) *adt.NodeWithParentPointer {
	
	if root == nil {
		return nil
	}

	depthFirst := Depth(first)
	depthSecond := Depth(second)

	if depthFirst > depthSecond {
		diff := depthFirst - depthSecond
		for diff > 0 {
			first = first.Parent
			diff--
			depthFirst--
		}
		root = first
	}

	if depthFirst < depthSecond {
		diff := depthSecond - depthFirst
		for diff > 0 {
			second = second.Parent
			diff--
			depthSecond--
		}
		root = second
	}

	if depthFirst == depthSecond {

		for first.Parent != second.Parent{
			first = first.Parent
			second = second.Parent
		}
	}

	return first 
}

func Depth(root *adt.NodeWithParentPointer) int {

	if root == nil {
		return 0
	}

	return Depth(root.Parent) + 1
}
	