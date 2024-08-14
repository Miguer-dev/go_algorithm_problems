package main

type LIFO []interface{}

func (s *LIFO) push(element interface{}) {
	*s = append(*s, element)
}

func (s *LIFO) pop() interface{} {
	deleted := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return deleted
}
