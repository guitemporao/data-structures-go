package queues

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(i int) {
	q.Items = append(q.Items, i)
}

func (q *Queue) Dequeue() int {
	lastIndex := len(q.Items) - 1
	toRemove := q.Items[lastIndex]
	q.Items = q.Items[1:]
	return toRemove
}

func (q *Queue) Peek() int {
	return q.Items[0]
}