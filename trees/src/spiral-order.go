package src

import (
	"bst/adt"
	"fmt"
)

func SpiralOrderTraversal(root *adt.Node) {
	if root == nil {
		return
	}
	treeHeight := treeHeight(root)

	rightToLeft := true
	for i:= 1; i <= treeHeight; i++ {
		printLevel(root, i, rightToLeft)
		rightToLeft = !rightToLeft
	}

}

func printLevel(root *adt.Node, level int, rightToLeft bool){
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Println(root)
		return
	}

	if rightToLeft {
		printLevel(root.Left, level - 1, rightToLeft)
		printLevel(root.Right, level - 1, rightToLeft)
		return
	}
	
	printLevel(root.Right, level - 1, rightToLeft)
	printLevel(root.Left, level - 1, rightToLeft)
	

}



