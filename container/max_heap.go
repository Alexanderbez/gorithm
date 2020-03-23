package container

type MaxHeap struct {
	Nodes []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		Nodes: make([]int, 0),
	}
}

// Insert inserts an integer into the max heap maintaining the max heap invariant
// that the root node is the maximum element.
func (mh *MaxHeap) Insert(x int) {
	mh.Nodes = append(mh.Nodes, x)

	i := len(mh.Nodes) - 1
	p := parentIndex(i)

	// bubble new node up as necessary to maintain max heap invariant
	for p >= 0 {
		if mh.Nodes[i] > mh.Nodes[p] {
			// swap and repeat
			mh.swap(p, i)
			i = p
			p = parentIndex(i)
		} else {
			// max heap invariant is maintained
			return
		}
	}
}

// Pop removes the max element (root node) from the max heap. If the heap is
// empty, (0, false) is returned.
func (mh *MaxHeap) Pop() (int, bool) {
	if len(mh.Nodes) == 0 {
		return 0, false
	}

	// get max (root) and replace root with the last element
	max := mh.Nodes[0]
	last := mh.Nodes[len(mh.Nodes)-1]
	mh.Nodes[0] = last
	mh.Nodes = mh.Nodes[:len(mh.Nodes)-1]

	var done bool

	s := len(mh.Nodes)
	i := 0

	// bubble down root node as necessary to maintain max heap invaraint
	for i < len(mh.Nodes) && !done {
		node := mh.Nodes[i]
		lIdx := leftChildIndex(i)
		rIdx := rightChildIndex(i)

		switch {
		case lIdx < s && mh.Nodes[lIdx] > node && rIdx < s && mh.Nodes[rIdx] > node:
			// Both children are greater than the current node, find max between children,
			// swap and repeat.
			if mh.Nodes[lIdx] > mh.Nodes[rIdx] {
				// swap with left child and repeat
				mh.swap(lIdx, i)
				i = lIdx
			} else {
				// swap with right child and repeat
				mh.swap(rIdx, i)
				i = rIdx
			}

		case lIdx < s && mh.Nodes[lIdx] > node:
			// swap with left child and repeat
			mh.swap(lIdx, i)
			i = lIdx

		case rIdx < s && mh.Nodes[rIdx] > node:
			// swap with right child and repeat
			mh.swap(rIdx, i)
			i = rIdx

		default:
			done = true
		}
	}

	return max, true
}

// Peek returns the maximum value (root node) of the max heap. If the heap is
// empty, (0, false) is returned.
func (mh MaxHeap) Peek() (int, bool) {
	if len(mh.Nodes) == 0 {
		return 0, false
	}

	return mh.Nodes[0], true
}

func (mh *MaxHeap) swap(a, b int) {
	tmp := mh.Nodes[a]
	mh.Nodes[a] = mh.Nodes[b]
	mh.Nodes[b] = tmp
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return (i * 2) + 1
}

func rightChildIndex(i int) int {
	return (i * 2) + 2
}
