package bst

// Node defines a single vertex in a binary search tree.
type Node struct {
	value       int
	left, right *Node
}

func NewNode(v int) *Node {
	return &Node{value: v}
}

// func (n Node) Value() int        { return n.value }
// func (n Node) LeftChild() *Node  { return n.left }
// func (n Node) RightChild() *Node { return n.right }

// Tree provides an implementation of a binary search tree.
type Tree struct {
	root  *Node
	count int
}

func NewTree() *Tree {
	return &Tree{}
}

// Size returns the current size (number of nodes) in the binary search tree.
func (t *Tree) Size() int { return t.count }

// Add attempts to add a new node into the binary search tree. Duplicates are not
// allowed. When a new node is inserted relative to it's parent, the invariant
// holds that it will be added to the left subtree if it is less than it's parent
// or added to the right subtree if it is greater than it's parent.
func (t *Tree) Add(v int) {
	newNode := NewNode(v)

	if t.root == nil {
		t.root = newNode
		t.count++
		return
	}

	var prev *Node
	curr := t.root

	for curr != nil {
		// no not allow duplicates
		if curr.value == v {
			return
		}

		prev = curr
		if v < curr.value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if v < prev.value {
		prev.left = newNode
	} else {
		prev.right = newNode
	}

	t.count++
}

// Search performs an iterative binary search on binary search tree for a given
// value. If the value exists in the tree, the corresponding Node object will be
// returned.
func (t *Tree) Search(v int) *Node {
	curr := t.root

	for curr != nil {
		if curr.value == v {
			return curr
		}

		if v < curr.value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return nil
}

// Traverse performs an in-order traversal of the binary search tree, where for
// each node, a given callback will be invoked with the node's value.
func (t *Tree) Traverse(cb func(i int)) {
	traverse(t.root, cb)
}

func traverse(n *Node, cb func(i int)) {
	if n == nil {
		return
	}

	traverse(n.left, cb)
	cb(n.value)
	traverse(n.right, cb)
}
