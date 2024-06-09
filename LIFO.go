package main

type LIFO struct {
	queue []interface{}
}

func (q *LIFO) push(element interface{}) {
	q.queue = append(q.queue, element)
}

func (q *LIFO) pop() interface{} {
	deleted := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]

	return deleted
}
