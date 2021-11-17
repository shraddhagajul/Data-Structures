package src

import (
	"bst/adt"
	"fmt"
)

func LeftView(root *adt.Node, current int ,maxLevelReached *int) {
	
	if root == nil {
		return
	}


	if *maxLevelReached < current {
		fmt.Println(root)

		*maxLevelReached = current
	}
		LeftView(root.Left,current + 1, maxLevelReached)
		LeftView(root.Right,current + 1 ,maxLevelReached)

}


func RightView(root *adt.Node, current int ,maxLevelReached *int) {
	
	if root == nil {
		return
	}
	if *maxLevelReached < current {
		fmt.Println(root)
		*maxLevelReached = current
	}

	RightView(root.Right,current + 1 ,maxLevelReached)
	RightView(root.Left,current + 1, maxLevelReached)
	

}