package queues

// FIRST IN FIRST OUT (FIFO)
type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(i int) {
	q.Items = append(q.Items, i)
}

func (q *Queue) Dequeue() int {
	lastIndex := len(q.Items) - 1 		// get the index of the last element
	toRemove := q.Items[lastIndex]	    // get the last element
	q.Items = q.Items[:lastIndex]       // remove the last element
	return toRemove                     // return the removed element
}

func (q *Queue) Peek() int {
	return q.Items[0]                  // get the first element
}	