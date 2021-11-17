package src

import "bst/adt"

func BstLCA(root *adt.Node, small int, big int) *adt.Node {
	if root == nil{
		return nil
	}

	if root.Val == small || root.Val == big {
		return root
	}
	if ( root.Val > small && root.Val < big ) {
		return root
	}

	if root.Val > small && root.Val > big {
		return BstLCA(root.Left,small,big)
	}
	
	return BstLCA(root.Right,small,big)
}