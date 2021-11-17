package main

import (
	"bst/adt"
	"fmt"
	"bst/src"
)

func main() {
  node := &adt.Node{10,nil,nil}
	node.Left = &adt.Node{5,nil,nil}
	node.Right = &adt.Node{20,nil,nil}
	node.Left.Left = &adt.Node{4,nil,nil}
	node.Left.Right = &adt.Node{6,nil,nil}
	node.Right.Left = &adt.Node{15,nil,nil}
	node.Right.Right = &adt.Node{22,nil,nil}


	src.BstInsertion(node,7)
	inorder(node)
	


	element := src.BstSearching(node,21)
	if element != nil{
		fmt.Println("Element : ", element.Val)

	}else{
		fmt.Println("Not found")
	}

	lca := src.BstLCA(node,5,22)
	fmt.Println("LCA : ",lca)

	kNode := src.KNodeBst(node, 7)
	if kNode != nil {
		fmt.Println("7th node : ", kNode)
	}

	diameter := src.DiameterOfTree(node)
	fmt.Println("Diameter : ",diameter)
  
	node1 := &adt.NodeWithParentPointer{10,nil,nil,nil}
	node1.Left = &adt.NodeWithParentPointer{4,nil,nil,node1}
	node1.Right = &adt.NodeWithParentPointer{20,nil,nil,node1}
	node1.Left.Left = &adt.NodeWithParentPointer{3,nil,nil,node1.Left}
	node1.Left.Right = &adt.NodeWithParentPointer{6,nil,nil,node1.Left}

	node1.Left.Right.Left = &adt.NodeWithParentPointer{5,nil,nil,node1.Left.Right}

	node1.Right.Left = &adt.NodeWithParentPointer{15,nil,nil,node1.Right}
	node1.Right.Right = &adt.NodeWithParentPointer{22,nil,nil,node1.Right}
	node1.Right.Right.Right = &adt.NodeWithParentPointer{23,nil,nil,node1.Right.Right}

	 fmt.Println("Node1 : ",node1)

	inorderParent(node1)
	successor := src.Successor(node1, 23)
	fmt.Println("Successor : ",successor)

	src.LevelOrderTraversal(node)

	src.SpiralOrderTraversal(node)

	isBst := src.IsBst(node)
	fmt.Println("Is bst" ,isBst)

	// dont want change any value in reursion pass pointer rather than value 
	level := 0
  src.LeftView(node,1,&level)
	level1 := 0
	src.RightView(node,1,&level1)

	src.LevelOrderUsingQueue(node)
	
	lcaWithParentPointer := src.LcaInBinaryTree(node1,node1,node1.Left.Right.Left,nil)
	fmt.Println("LCA : ",lcaWithParentPointer)
}


func inorder(node *adt.Node) {
	
	if node == nil{
		return
	}

	inorder(node.Left)
	fmt.Println(node)
	inorder(node.Right)
}

func inorderParent(node *adt.NodeWithParentPointer) {
	
	if node == nil{
		return
	}

	inorderParent(node.Left)
	fmt.Println(node)
	inorderParent(node.Right)
}