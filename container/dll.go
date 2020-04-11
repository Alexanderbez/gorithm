package container

type Node struct {
	Value interface{}

	list *DoubleLinkedList
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	root *Node
	len  int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	l := &DoubleLinkedList{
		root: &Node{},
	}

	l.root.prev = l.root
	l.root.next = l.root

	return l
}

func (l *DoubleLinkedList) Len() int {
	return l.len
}

func (l *DoubleLinkedList) Front() *Node {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *DoubleLinkedList) Back() *Node {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

func (l *DoubleLinkedList) InsertFront(v interface{}) *Node {
	n := &Node{Value: v}
	l.insertAfter(n, l.root)
	return n
}

func (l *DoubleLinkedList) InsertBack(v interface{}) *Node {
	n := &Node{Value: v}
	l.insertAfter(n, l.root.prev)
	return n
}

func (l *DoubleLinkedList) Remove(n *Node) interface{} {
	if n == nil {
		return nil
	}

	if n.list == l {
		l.delink(n)
		n.list = nil
		l.len--
	}

	return n.Value
}

func (l *DoubleLinkedList) MoveToFront(n *Node) {
	l.delink(n)
	l.moveAfter(n, l.root)
}

func (l *DoubleLinkedList) MoveToBack(n *Node) {
	l.delink(n)
	l.moveAfter(n, l.root.prev)
}

func (l *DoubleLinkedList) insertAfter(n, at *Node) {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
}

func (l *DoubleLinkedList) delink(n *Node) {
	if n == nil || n.list != l {
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
}

func (l *DoubleLinkedList) moveAfter(n, at *Node) {
	if n == nil || n.list != l {
		return
	}

	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
}
