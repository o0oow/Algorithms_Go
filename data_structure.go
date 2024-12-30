package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val}
}
