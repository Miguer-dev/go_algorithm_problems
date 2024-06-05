package main

func isValid(s string) bool {
	queue := LIFO{}
	openClose := make(map[rune]rune)
	openClose[')'] = '('
	openClose[']'] = '['
	openClose['}'] = '{'

	for _, value := range s {
		if _, exits := openClose[value]; !exits {
			//if ( o [ o {
			queue.push(value)
		} else if len(queue.queue) != 0 {
			//if ) o ] o }
			open := queue.pop()
			if open != openClose[value] {
				return false
			}
		} else {
			return false
		}
	}

	if len(queue.queue) != 0 {
		return false
	}

	return true
}

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
