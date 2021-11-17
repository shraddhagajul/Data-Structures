package adt

type Node struct{
	Val int
	Left *Node
	Right *Node
}

type NodeWithParentPointer struct {
	Val int
	Left *NodeWithParentPointer
	Right *NodeWithParentPointer
	Parent *NodeWithParentPointer
}