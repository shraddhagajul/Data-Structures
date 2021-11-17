package src

import (
	"bst/adt"
	"fmt"
)

func DiameterOfTree(root *adt.Node) int{

	leftSubtreeHeight := heightOfNode(root.Left)
	fmt.Println("Height of left subtree : ",leftSubtreeHeight)

	rightSubtreeHeight := heightOfNode(root.Right)
	fmt.Println("Height of right subtree : ",rightSubtreeHeight)

	return leftSubtreeHeight + 1 + rightSubtreeHeight 
}

func heightOfNode(root *adt.Node) int {
	fmt.Println("Root : ",root)
	if root == nil {
		return 0
	}

	lHeight := heightOfNode(root.Left)
	rHeight := heightOfNode(root.Right)
	fmt.Println("lheight ", lHeight )
	fmt.Println("rheight ", rHeight )
	if lHeight > rHeight {
		
		return lHeight + 1
	}
	
	return rHeight + 1
	
}