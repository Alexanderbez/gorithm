package container

// Queue defines a queue data structure providing FIFO semantics, where each
// item in the queue is an interface.
type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

// Push pushes an item onto the queue.
func (q *Queue) Push(i interface{}) {
	q.items = append([]interface{}{i}, q.items...)
}

// Pop removes an item from the queue. If the queue is empty, nil is returned.
func (q *Queue) Pop() interface{} {
	if q.Size() == 0 {
		return nil
	}

	v := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]

	return v
}
