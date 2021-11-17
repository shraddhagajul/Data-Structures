package src

import( "bst/adt")

func BstInsertion(node *adt.Node, newElement int){
	if newElement < node.Val {
		if node.Left == nil {
			node.Left = &adt.Node{Val: newElement, Left: nil, Right: nil}
			return
		}

		BstInsertion(node.Left, newElement)
		return
	}

	if newElement > node.Val {
		if node.Right == nil {
			node.Right = &adt.Node{Val: newElement, Left: nil, Right: nil}
			return
		}

		BstInsertion(node.Right, newElement)
		
	}
}


