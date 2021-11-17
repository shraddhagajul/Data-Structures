package src

import (
	"bst/adt"

)

func KNodeBst(root *adt.Node, rank int) *adt.Node {

	if root == nil {
		return nil
	}

	leftSubtreeCount := count(root.Left)

	currentRank := leftSubtreeCount + 1

	if currentRank > rank {
		return KNodeBst(root.Left, rank)
	}

	if currentRank < rank {
		return KNodeBst(root.Right, rank - currentRank)
	}

	return root
}

func count(root *adt.Node) int{
	if root == nil{
		return 0
	}
	return count(root.Left) + 1 + count(root.Right)
}