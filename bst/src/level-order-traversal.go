package src

import (
	"bst/adt"
	"fmt"
)

func LevelOrderTraversal(root *adt.Node) {
	if root == nil {
		return
	}
	treeHeight := treeHeight(root)

	for i:= 1; i <= treeHeight; i++ {
		printCurrentLevel(root, i)
	}

	
}

func printCurrentLevel(root *adt.Node,level int) {
	
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Println(root)
		return
	}
	
	if level > 1 {
		fmt.Println("Root.left :" ,root.Left, level - 1)
		printCurrentLevel(root.Left,level - 1)
		fmt.Println("Root.Right :" ,root.Right , level -1)
		printCurrentLevel(root.Right, level - 1)
	}

}


func treeHeight(root *adt.Node) int {
	if root == nil {
		return 0
	}
	lHeight := treeHeight(root.Left)
	rHeight := treeHeight(root.Right)

	if lHeight > rHeight {
		return lHeight + 1
	}
	
	return rHeight + 1

}

