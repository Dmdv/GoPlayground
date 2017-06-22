package singlelinkedlist

// Next returns next
func (node *Node) Next() *Node {
	if node == nil {
		return nil
	}
	return node.next
}

// InsertToHead inserts to head
func (list *Singlelinkedlist) InsertToHead(node *Node) *Node {

	list.len++

	if list.head == nil {
		list.head = node
		return node
	}

	node.next = list.head
	list.head = node

	return node
}

// InsertToTail inserts to tail
func (list *Singlelinkedlist) InsertToTail(node *Node) *Node {

	list.len++

	if list.head == nil {
		list.head = node
		return node
	}

	tail := list.head

	for tail.next != nil {
		tail = tail.next
	}

	tail.next = node
	return node
}

// Remove removes a node from list
func (list *Singlelinkedlist) Remove(node *Node) *Node {

	// check
	if list == nil || list.head == nil || node == nil {
		return nil
	}

	// if this is the firt node - compare and remove
	if list.head == node {
		list.head = list.head.next
		node.next = nil
		list.len--
		return node
	}

	// iterate through all nodes
	for current := list.head; current != nil; current = current.next {
		if current.next == node {
			list.len--
			current.next = current.next.next
			node.next = nil
			return node
		}
	}

	return nil
}

// First, Last, Len, ToSlice, NewFromSlice

// First returns head
func (list *Singlelinkedlist) First() *Node {
	if list == nil {
		return nil
	}
	return list.head
}

// Last returns tail
func (list *Singlelinkedlist) Last() *Node {

	if list == nil {
		return nil
	}

	node := list.head

	for node != nil {
		node = node.next
	}

	return node
}

// Len return a length of list
func (list *Singlelinkedlist) Len() int {
	if list == nil {
		return 0
	}

	return list.len
}

// ToSlice convert list to slice
func (list *Singlelinkedlist) ToSlice() (s []interface{}) {
	for p := list.head; p != nil; p = p.next {
		s = append(s, p.Data)
	}
	return s
}

// NewFromSlice return new list from slice
func NewFromSlice(data []interface{}) *Singlelinkedlist {
	list := new(Singlelinkedlist)

	for _, n := range data {
		list.InsertToTail(&Node{Data: n})
	}

	return list
}
