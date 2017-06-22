package singlelinkedlist

// Node is a node of a linked list.
type Node struct {
	Data interface{}
	next *Node
}

// BiNode is a Double linked node
type BiNode struct {
	Data interface{}
	next *Node
	prev *Node
}

// Singlelinkedlist singled linked list
type Singlelinkedlist struct {
	head *Node
	len  int
}
